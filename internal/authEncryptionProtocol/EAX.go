package authEncryptionProtocol

import (
	"infbez_labs/internal/alphabet"
)

type EAX struct {
	CFB *CFB
}

func NewEAX(alphabet *alphabet.Alphabet) *EAX {
	cfb := NewCFB(alphabet)

	return &EAX{CFB: cfb}
}

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
	data = append(data, []rune("_____")...)

	cMAC := e.CFB.Forward(data, []rune(secIn), spNetSeed, -1)
	cIV := []rune(e.CFB.Forward(append([]rune(secIn), tmp...), packet.IV, spNetSeed, -1))
	if mac == 1 {
		tmp = []rune(e.CFB.Forward(packet.Message, cIV, spNetSeed, -1))
		tmp = e.CFB.alphabet.BlockXOR(tmp, cIV)
		MAC = e.CFB.alphabet.BlockXOR(tmp, []rune(cMAC))
		MAC = e.CFB.alphabet.BlockXOR(packet.MAC, []rune(cMAC))
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
		IV:        cIV,
		Message:   MSG,
		MAC:       MAC,
	}
}
