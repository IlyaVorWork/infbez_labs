package trithemius_test

import (
	l1 "infbez_labs/Lab1"
	"testing"
)

var TelegraphAlphabet = []rune("АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЫЬЭЮЯ_")

func TestSTrithemius_nearbyEntrances(t *testing.T) {
	alphabet := l1.NewTrithemius(TelegraphAlphabet)

	var (
		IN1 = "ОРЕХ"
		IN2 = "ОПЕХ"
		IN3 = "ОПЕФ"
	)

	var (
		KEY1 = "ХОРОШО_БЫТЬ_ВАМИ"
		KEY2 = "МОЛЧАНИЕ_ЗОЛОТО_"
	)

	tests := []struct {
		name       string
		openText   string
		key        string
		cipherText string
	}{
		{"ОРЕХ_ХорошоБытьВами", IN1, KEY1, "ТЬЧЫ"},
		{"ОПЕХ_ХорошоБытьВами", IN2, KEY1, "ТАЧЫ"},
		{"ОПЕФ_ХорошоБытьВами", IN3, KEY1, "ТАЧС"},
		{"ОРЕХ_МолчаниеЗолото", IN1, KEY2, "ЗЖБИ"},
		{"ОПЕХ_МолчаниеЗолото", IN2, KEY2, "ЗДБИ"},
		{"ОПЕФ_МолчаниеЗолото", IN3, KEY2, "ЗДБЕ"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alphabet.EncodeSTrithemius(tt.openText, tt.key)

			if got != tt.cipherText {
				t.Errorf("Thrithemus(text=%q , key=%q), want %q but return %q", tt.openText, tt.key, tt.cipherText, got)
				return
			}

		})
	}
}

func TestAlphabet(t *testing.T) {
	alphabet := l1.NewTrithemius(TelegraphAlphabet)

	var (
		K1 = "ДИНОЗАВР_ЗАУРОПОД"
		K2 = "ГАМЕЛЬНСКИЙ_АНТИКВАР"
		K3 = "ГАРРИ_ПОТТЕР_И_ФИЛОСОФСКИЙ_КАМЕНЬ"

		OutputTable1 = "ДИНОЗАВР_ЙБУСПТФЕГЖКЛМХЦЧШЩЫЬЭЮЯ"
		OutputTable2 = "ГАМЕЛЬНСКИЙ_БОТПРВДУЖЗФХЦЧШЩЫЭЮЯ"
		OutputTable3 = "ГАРСИ_ПОТУЕФБЙВХКЛЦЧШЩЫМНЬДЭЖЮЗЯ"
	)

	tests := []struct {
		name        string
		inputTable  string
		outputTable string
	}{
		{"Создание Таблицы. Проверка 1", K1, OutputTable1},
		{"Создание Таблицы. Проверка 2", K2, OutputTable2},
		{"Создание Таблицы. Проверка 3", K3, OutputTable3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := string(alphabet.BuildTrithemiusAlphabet(tt.inputTable))

			if tt.outputTable != got {
				t.Errorf("Faild BuildTrithemiusAlphabet(input=%q), want %v but return %v", tt.inputTable, tt.outputTable, got)
				return
			}

		})
	}
}
