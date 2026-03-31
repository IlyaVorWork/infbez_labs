package authEncryptionProtocol

import (
	"errors"
	"fmt"
	"infbez_labs/internal/alphabet"
	"slices"
)

var (
	starter = []byte{1, 0, 0}
	ender   = []byte{0, 0, 1}
)

type ProtocolSimulator struct {
	alphabet *alphabet.Alphabet
	EAX      *EAX

	msgCnt, last    int
	iv0             []rune
	AssData         AssData
	spNetSeed       string
	dataMac, secret string
}

func NewProtocol(alphabet *alphabet.Alphabet) *ProtocolSimulator {
	eax := NewEAX(alphabet)
	return &ProtocolSimulator{
		alphabet: alphabet,
		EAX:      eax,
	}
}

func (p *ProtocolSimulator) NewConnection(assData AssData, keyIn, nonse string) {

	t1 := append(assData.Reciever[:], assData.Sender[:]...)
	t2 := append(assData.MType[:], assData.Transmission[:]...)
	t2 = append(t2, '_', '_', '_', '_', '_')
	t3 := make([]rune, 0, len(t1))
	if string(assData.Reciever[:]) < string(assData.Sender[:]) {
		t3 = t1
	} else {
		t3 = append(assData.Sender[:], assData.Reciever[:]...)
	}

	v0 := p.alphabet.AddTxt(string(t1), string(t2))
	v0 = p.alphabet.AddTxt(v0, nonse)
	iv0 := []rune(v0)[:12] ///

	secret := p.EAX.CFB.Forward(append(t3, t2...), []rune(keyIn), keyIn, -1)

	data := append(assData.MType[:], assData.Sender[:]...)
	data = append(data, assData.Reciever[:]...)
	data = append(data, assData.Transmission[:]...)
	data = append(data, '_', '_', '_', '_') // 4
	dataMac := p.EAX.CFB.Forward(data, []rune(secret), keyIn, -1)

	{
		p.msgCnt = -1
		p.last = -1
		p.iv0 = iv0
		p.AssData = assData
		p.spNetSeed = keyIn
		p.dataMac = dataMac
		p.secret = secret
	}
}

func (p *ProtocolSimulator) Send(message string) []byte {
	var out []byte
	var secPacket Packet
	p.msgCnt++
	iv := append(p.iv0, p.EAX.CFB.alphabet.NumToBlock(p.msgCnt)...)

	packet := p.PreparePacket(p.AssData, string(iv), message)

	switch p.AssData.MType {
	case [2]rune{'В', '_'}:
		secPacket = packet
		out = p.Transmit(packet)
	case [2]rune{'В', 'А'}:
		secPacket = p.EAX.ForwardEaxCfb(packet, p.dataMac, p.spNetSeed, p.secret, 1)
		out = p.Transmit(secPacket)
	case [2]rune{'В', 'Б'}:
		secPacket = p.EAX.ForwardEaxCfb(packet, p.dataMac, p.spNetSeed, p.secret, 0)
		out = p.Transmit(secPacket)
	}
	return out
}

func (p *ProtocolSimulator) Recieve(bits []byte) Packet {
	packet := p.Response(bits)
	current := p.EAX.CFB.alphabet.BlockToNum(packet.IV[12:16])
	var out Packet

	if current <= p.last {
		panic(errors.New("received packet какой-то странный"))
	}

	switch mType := packet.AssocData.MType; {
	case slices.Equal(mType[:], []rune{'В', 'Б'}):
		recPacker := p.EAX.InvertEaxCfb(packet, p.spNetSeed, p.secret, 0)
		recPacker.Message = p.UnpadMessage(recPacker.Message)

		if slices.Equal(recPacker.MAC, slices.Repeat([]rune{'_'}, 16)) {
			p.last = current
			recPacker.MAC = []rune("ОК")
		} else {
			fmt.Println("Блфть")
		}
		out = recPacker

	case slices.Equal(mType[:], []rune{'В', 'А'}) && !slices.Equal(p.AssData.MType[:], []rune{'В', 'Б'}):
		recPacker := p.EAX.InvertEaxCfb(packet, p.spNetSeed, p.secret, 1)
		recPacker.Message = p.UnpadMessage(recPacker.Message)

		if slices.Equal(recPacker.MAC, slices.Repeat([]rune{'_'}, 16)) {
			p.last = current
			recPacker.MAC = []rune("ОК")
		}
		out = recPacker

	case slices.Equal(mType[:], []rune{'В', '_'}) && slices.Equal(p.AssData.MType[:], []rune{'В', '_'}):
		recPacker := packet
		recPacker.Message = p.UnpadMessage(recPacker.Message)

		if slices.Equal(recPacker.MAC, []rune{}) {
			p.last = current
			recPacker.MAC = []rune("N/A")
		}
		out = recPacker

	default:
		out = packet
	}

	return out
}

func (p *ProtocolSimulator) Invert(stream []byte, n int) []byte {
	stream[n] = (stream[n] + 1) % 2
	return stream
}

func (p *ProtocolSimulator) Transmit(packet Packet) []byte {
	out := append([]rune{}, packet.AssocData.MType[:]...)
	out = append(out, packet.AssocData.Sender[:]...)
	out = append(out, packet.AssocData.Reciever[:]...)
	out = append(out, packet.AssocData.Transmission[:]...)
	out = append(out, packet.AssocData.Length[:]...)
	out = append(out, packet.IV...)
	out = append(out, packet.Message...)
	out = append(out, packet.MAC...)
	stream, _ := p.alphabet.MessageToBin(out)
	return stream
}

func (p *ProtocolSimulator) Response(stream []byte) Packet {
	streamRune := p.alphabet.BinToMessage(stream)
	mesLen := len(streamRune)

	packetType := streamRune[:2]
	packetSender := streamRune[2:10]
	packetReciever := streamRune[10:18]
	packetSession := streamRune[18:27]
	packetLength := streamRune[27:32]
	packetIV := streamRune[32:48]

	L := 0
	for i := 0; i < 5; i++ {
		key, _ := p.alphabet.GetKeyByChar(packetLength[i])
		L = 32*L + key
	}
	L /= 5
	message := streamRune[48 : 48+L]
	mac := streamRune[48+L : mesLen]

	packet := Packet{
		AssocData: AssData{
			[2]rune(packetType),
			[8]rune(packetSender),
			[8]rune(packetReciever),
			[9]rune(packetSession),
			[5]rune(packetLength),
		},
		IV:      packetIV,
		Message: message,
		MAC:     mac,
	}
	return packet
}

func (p *ProtocolSimulator) PreparePacket(dataIn AssData, IV, message string) Packet {
	iv := p.alphabet.AddTxt("________________", IV)
	msg := p.PadMessage([]rune(message))

	msgBin, err := p.alphabet.MessageToBin(msg)
	if err != nil {
		panic(err)
	}
	msgBinLen := len(msgBin)

	var a []rune
	for i := 0; i < 5; i++ {
		a = append([]rune{p.alphabet.GetCharByKey(msgBinLen % 32)}, a...)
		msgBinLen /= 32
	}
	dataIn.Length = [5]rune(a)
	mac := ""

	return Packet{
		AssocData: dataIn,
		IV:        []rune(iv),
		Message:   msg,
		MAC:       []rune(mac),
	}
}

func (p *ProtocolSimulator) checkPadding(binData []byte) (hasPad bool, numBlocks int, padLength int) {
	var (
		binDataLen = len(binData)
		remain     = binDataLen % 80
		blocks     = binDataLen / 80
	)

	if remain != 0 {
		return false, 0, 0
	}

	if !slices.Equal(binData[binDataLen-3:binDataLen], ender) {
		return false, 0, 0
	}

	NB := binData[binDataLen-13 : binDataLen-3]  // Блок с инфой кол-ва блоков
	PL := binData[binDataLen-20 : binDataLen-13] // Блок с инфой длины подложки

	for i := 0; i < 7; i++ {
		padLength *= 2
		padLength += int(PL[i])
	}

	for i := 0; i < 10; i++ {
		numBlocks *= 2
		numBlocks += int(NB[i])
	}

	if (numBlocks != blocks) || (padLength < 23) || (padLength >= 103) {
		return false, 0, 0
	}

	if !slices.Equal(binData[binDataLen-padLength:binDataLen-padLength+3], starter) {
		return false, 0, 0
	}

	for i := binDataLen - padLength + 3; i < binDataLen-20; i++ {
		if int(binData[i]) == 1 {
			return false, 0, 0
		}
	}

	return true, numBlocks, padLength
}

func (p *ProtocolSimulator) producePadding(remain int, numBlocks int) []byte {
	var (
		resultBlockNum int // Количество итоговых блоков
		resultPadLen   int // Места для подложки
		pad            []byte
	)

	if remain == 0 {
		resultBlockNum = numBlocks + 1
		resultPadLen = 80
	} else if remain <= 57 {
		resultBlockNum = numBlocks + 1
		resultPadLen = 80 - remain
	} else {
		resultBlockNum = numBlocks + 2
		resultPadLen = 160 - remain
	}

	pad = make([]byte, 0, resultPadLen)
	pad = append(pad, starter...)
	pad = append(pad, slices.Repeat([]byte{0}, resultPadLen-23)...)

	for i := 6; i >= 0; i-- {
		pad = append(pad, byte((resultPadLen>>i)&1))
	}

	for i := 9; i >= 0; i-- {
		pad = append(pad, byte((resultBlockNum>>i)&1))
	}

	pad = append(pad, ender...)
	return pad
}

func (p *ProtocolSimulator) PadMessage(message []rune) []rune {
	bins, err := p.alphabet.MessageToBin(message)
	if err != nil {
		panic(err)
	}

	binLen := len(bins)
	blockNum := binLen / 80
	remainder := binLen % 80
	hasPad := true

	if remainder == 0 {
		hasPad, _, _ = p.checkPadding(bins)
	}

	if hasPad {
		pad := p.producePadding(remainder, blockNum)
		bins = append(bins, pad...)
	}

	return p.alphabet.BinToMessage(bins)
}

func (p *ProtocolSimulator) UnpadMessage(message []rune) []rune {
	var (
		bins, err         = p.alphabet.MessageToBin(message)
		binsLen           = len(bins)
		hasPad, _, padLen = p.checkPadding(bins)
		resultRuneMessage = message
	)

	if err != nil {
		panic(err)
	}

	if hasPad {
		resultBinsMessage := bins[0 : binsLen-padLen]
		resultRuneMessage = p.alphabet.BinToMessage(resultBinsMessage)
	}

	return resultRuneMessage
}
