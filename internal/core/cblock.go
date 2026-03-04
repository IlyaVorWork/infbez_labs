package core

var innerState = [4]string{
	"________________",
	"ПРОЖЕКТОР_ЧЕПУХИ",
	"КОЛЫХАТЬ_ПАРОДИЮ",
	"КАРМАННЫЙ_АТАМАН",
}

type CBlock struct {
	state [4]string
}

func NewCBlock() *CBlock {
	return &CBlock{
		state: innerState,
	}
}
