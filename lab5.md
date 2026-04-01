
# KDF

### KDFRun
```
    func (k KeyDerivation) Run(input, salt string, context []string, outputSizes []int, rounds int) []string {
	tmp := input + salt

	for i := 0; i <= rounds; i++ {
		ext := k.hasher.Hash(tmp)
		tmp = ext + tmp
	}

	prk := tmp

	var out []string

	for i := 0; i < len(outputSizes); i++ {
		q := (outputSizes[i] - (outputSizes[i] % 64)) / 64

		rem := i
		res := ""

		for rem > 0 {
			h := rem % 32
			res = res + string(k.alphabet.GetCharByKey(h))
			rem = (rem - h) / 32
		}

		if q > 0 {
			hash2 := prk

			for j := 0; j <= q; j++ {
				tmp = hash2 + context[i] + prk
				hash2 = k.hasher.Hash(tmp)
				res = hash2 + res
			}
		} else {
			tmp = prk + context[i] + prk
			res = k.hasher.Hash(tmp)
		}

		out = append(out, string([]rune(res)[:outputSizes[i]]))
	}

	return out
}
```

# ProtocolSimulator

### PadMessage

```
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
```

### UnpadMessage

```
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
```

### PreparePacket

```
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
```

### Transmit

```
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
```

### Response

```
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
```

# CFB

### CFB_Forward

```
func (c *CFB) Forward(message, iv []rune, spNetSeed string, macIn int8) string {
	var (
		ground    = []rune("________________") // 16
		out       []rune
		keyStream []rune
	)
	blockNum := len(message) / blockSize

	for i := 0; i < blockNum; i++ {
		inp := message[16*i : (i+1)*16]
		ground = c.alphabet.BlockXOR(inp, ground)
		keyStream = []rune(c.spNet.FrwSPNet(string(iv), spNetSeed, rounds))
		iv = c.alphabet.BlockXOR(inp, keyStream)
		out = append(out, iv...)
	}

	keyStream = []rune(c.spNet.FrwSPNet(string(iv), spNetSeed, rounds))
	mac := c.alphabet.BlockXOR(ground, keyStream)
	if macIn == 1 {
		out = append(out, mac...)
	} else if macIn == -1 {
		out = mac
	}
	return string(out)
}
```

### CFB_Invert

```
func (c *CFB) Invert(message, iv []rune, spNetSeed string, macIn int8) string {
	var (
		ground     = []rune("________________") // 16
		out        []rune
		keyStream  []rune
		msgLen     = len(message) / blockSize
		dataBlocks = msgLen
	)
	if macIn != 0 {
		if msgLen == 0 {
			return errors.New("ошибка входных данных: сообщение должно быть не короче 16 символов").Error()
		}
		dataBlocks = msgLen - 1
	}

	for i := 0; i < dataBlocks; i++ {
		cipherBlock := message[blockSize*i : (i+1)*blockSize]
		keyStream = []rune(c.spNet.FrwSPNet(string(iv), spNetSeed, rounds))
		iv = cipherBlock
		plaintBlock := c.alphabet.BlockXOR(cipherBlock, keyStream)
		ground = c.alphabet.BlockXOR(ground, plaintBlock)
		out = append(out, plaintBlock...)

	}

	if macIn != 0 {
		recvMAC := message[(msgLen-1)*blockSize:]
		keyStream = []rune(c.spNet.FrwSPNet(string(iv), spNetSeed, rounds))
		keyStream = c.alphabet.BlockXOR(recvMAC, keyStream)
		calcMAC := c.alphabet.BlockXOR(ground, keyStream)

		if macIn == 1 {
			out = append(out, calcMAC...)
		} else {
			out = calcMAC
		}
	}
	return string(out)
}
```

# EAX

### EAX_Forward

```
func (e *EAX) ForwardEaxCfb(packet Packet, cMacIn, spNetSeed, secIn string, mac int8) Packet {
	tmp := make([]rune, 0, 15)
	var MSG []rune
	var MAC []rune

	tmp = append(tmp, packet.AssocData.MType[:]...)
	tmp = append(tmp, packet.AssocData.Transmission[:]...)
	tmp = append(tmp, packet.AssocData.Length[:]...)

	cIV := e.CFB.Forward(append([]rune(secIn), tmp...), packet.IV, spNetSeed, -1)
	if mac == 1 {
		tmp = []rune(e.CFB.Forward(packet.Message, []rune(cIV), spNetSeed, -1))
		MAC = e.CFB.alphabet.BlockXOR(tmp, []rune(cIV))
		MAC = e.CFB.alphabet.BlockXOR(MAC, []rune(cMacIn))
		MSG = packet.Message
	} else {
		tmp = []rune(e.CFB.Forward(packet.Message, []rune(cIV), spNetSeed, 1))
		m := tmp[len(packet.Message) : len(packet.Message)+16]
		MAC = e.CFB.alphabet.BlockXOR(m, []rune(cIV))
		MAC = e.CFB.alphabet.BlockXOR(MAC, []rune(cMacIn))
		MSG = tmp[:len(packet.Message)]
	}

	return Packet{
		AssocData: packet.AssocData,
		IV:        packet.IV,
		Message:   MSG,
		MAC:       MAC,
	}
}
```

### EAX_Invert

```
func (e *EAX) InvertEaxCfb(packet Packet, spNetSeed, secIn string, mac int8) Packet {
	tmp := make([]rune, 0, 15)
	data := make([]rune, 0, 15)
	var MSG []rune
	var MAC []rune

	tmp = append(tmp, packet.AssocData.MType[:]...)
	tmp = append(tmp, packet.AssocData.Transmission[:]...)
	tmp = append(tmp, packet.AssocData.Length[:]...)

	data = append(data, packet.AssocData.MType[:]...)
	data = append(data, packet.AssocData.Sender[:]...)
	data = append(data, packet.AssocData.Reciever[:]...)
	data = append(data, packet.AssocData.Transmission[:]...)
	data = append(data, []rune("____")...)

	cMAC := e.CFB.Forward(data, []rune(secIn), spNetSeed, -1)
	cIV := []rune(e.CFB.Forward(append([]rune(secIn), tmp...), packet.IV, spNetSeed, -1))
	if mac == 1 {
		tmp = []rune(e.CFB.Forward(packet.Message, cIV, spNetSeed, -1))
		tmp = e.CFB.alphabet.BlockXOR(tmp, cIV)
		MAC = e.CFB.alphabet.BlockXOR(tmp, []rune(cMAC))
		MAC = e.CFB.alphabet.BlockXOR(packet.MAC, MAC)
		MSG = packet.Message
	} else {
		cont := e.CFB.alphabet.BlockXOR(packet.MAC, cIV)
		cont = e.CFB.alphabet.BlockXOR(cont, []rune(cMAC))

		tmp = []rune(e.CFB.Invert(append(packet.Message, cont...), cIV, spNetSeed, 1))
		m := tmp[len(packet.Message) : len(packet.Message)+16]

		MAC = m
		MSG = tmp[:len(packet.Message)]
	}

	return Packet{
		AssocData: packet.AssocData,
		IV:        packet.IV,
		Message:   MSG,
		MAC:       MAC,
	}
}
```

# EAX_CFB

### EAX_CFB (так и назови)

```
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
```