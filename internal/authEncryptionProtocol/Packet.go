package authEncryptionProtocol

import (
	"fmt"
	"infbez_labs/internal/alphabet"
)

type AssData struct {
	MType        [2]rune
	Sender       [8]rune
	Reciever     [8]rune
	Transmission [9]rune
	Length       [5]rune
}

func NewAssData(mType, sender, reciever, transmission, length string) AssData {
	return AssData{
		MType:        [2]rune([]rune(mType)),
		Sender:       [8]rune([]rune(sender)),
		Reciever:     [8]rune([]rune(reciever)),
		Transmission: [9]rune([]rune(transmission)),
		Length:       [5]rune([]rune(length)),
	}
}

type Packet struct {
	AssocData AssData
	IV        []rune
	Message   []rune
	MAC       []rune
}

func NewPacket(alpha *alphabet.Alphabet) *ProtocolSimulator {
	return &ProtocolSimulator{alphabet: alpha}
}

func (pk *Packet) String() string {
	return fmt.Sprintf(
		"\tType = %s(%d)\n\tSender = %s(%d)\n\tReceiver = %s(%d)\n\tSession = %s(%d)\n\tLength = %s(%d)\nIV = %s(%d)\nMessage = %s(%d)\nMAC = %s(%d)",
		string(pk.AssocData.MType[:]),
		len(pk.AssocData.MType),
		string(pk.AssocData.Sender[:]),
		len(pk.AssocData.Sender),
		string(pk.AssocData.Reciever[:]),
		len(pk.AssocData.Reciever),
		string(pk.AssocData.Transmission[:]),
		len(pk.AssocData.Transmission),
		string(pk.AssocData.Length[:]),
		len(pk.AssocData.Length),
		string(pk.IV),
		len(pk.IV),
		string(pk.Message),
		len(pk.Message),
		string(pk.MAC),
		len(pk.MAC),
	)
}

func (ad *AssData) String() string {
	return fmt.Sprintf(
		"\tType = %s \n\tSender = %s\n\tReceiver = %s\n\tSession = %s\n\tLength = %s",
		string(ad.MType[:]),
		string(ad.Sender[:]),
		string(ad.Reciever[:]),
		string(ad.Transmission[:]),
		string(ad.Length[:]),
	)
}
