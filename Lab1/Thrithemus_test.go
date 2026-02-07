package Thrithemus_test

import (
	/* "fmt" */
	l1 "infbez_labs/Lab1"
	"testing"
)

func TestSThrithemus_nearbyEntrances(t *testing.T) {
	alphabet := l1.NewAlphabet()

	var (
		IN1 string = "ОРЕХ"
		IN2 string = "ОПЕХ"
		IN3 string = "ОПЕФ"
	)

	var (
		KEY1 string = "ХОРОШО_БЫТЬ_ВАМИ"
		KEY2 string = "МОЛЧАНИЕ_ЗОЛОТО_"
	)

	tests := []struct {
		name       string
		openText   string
		key        string
		chiperText string
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
			got := alphabet.EncodeSThrithemus(tt.openText, tt.key)

			if got != tt.chiperText {
				t.Errorf("Thrithemus(text=%q , key=%q), want %q but return %q", tt.openText, tt.key, tt.chiperText, got)
				return
			}

		})
	}
}

func TestAlphabet(t *testing.T) {
	alphabet := l1.NewAlphabet()

	var (
		K1 string = "ДИНОЗАВР_ЗАУРОПОД"
		K2 string = "ГАМЕЛЬНСКИЙ_АНТИКВАР"
		K3 string = "ГАРРИ_ПОТТЕР_И_ФИЛОСОФСКИЙ_КАМЕНЬ"

		OutputTable1 string = "ДИНОЗАВР_ЙБУСПТФЕГЖКЛМХЦЧШЩЫЬЭЮЯ"
		OutputTable2 string = "ГАМЕЛЬНСКИЙ_БОТПРВДУЖЗФХЦЧШЩЫЭЮЯ"
		OutputTable3 string = "ГАРСИ_ПОТУЕФБЙВХКЛЦЧШЩЫМНЬДЭЖЮЗЯ"
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
			got := string(alphabet.BuildThrithemusAlphabet(tt.inputTable))

			if tt.outputTable != got {
				t.Errorf("Faild BuildThrithemusAlphabet(input=%q), want %v but return %v", tt.inputTable, tt.outputTable, got)
				return
			}

		})
	}
}
