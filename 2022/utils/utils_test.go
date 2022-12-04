package utils

import (
	"testing"
)

func TestToClipboard(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		ToClipboard("test")
	})
}

func TestToInt(t *testing.T) {
	t.Run("string_1", func(t *testing.T) {
		if ToInt("1") != 1 {
			t.Errorf("ToInt(\"1\") != 1")
		}
	})
	t.Run("string_2", func(t *testing.T) {
		if ToInt("2") != 2 {
			t.Errorf("ToInt(\"2\") != 2")
		}
	})
	t.Run("string_123", func(t *testing.T) {
		if ToInt("123") != 123 {
			t.Errorf("ToInt(\"123\") != 123")
		}
	})
}

func TestToString(t *testing.T) {
	t.Run("int_1", func(t *testing.T) {
		if ToString(1) != "1" {
			t.Errorf("ToString(1) != \"1\"")
		}
	})
	t.Run("int_2", func(t *testing.T) {
		if ToString(2) != "2" {
			t.Errorf("ToString(2) != \"2\"")
		}
	})
	t.Run("int_123", func(t *testing.T) {
		if ToString(123) != "123" {
			t.Errorf("ToString(123) != \"123\"")
		}
	})
	t.Run("byte", func(t *testing.T) {
		if got := ToString(byte(65)); got != "A" {
			t.Errorf("ToString(byte(65)) != \"A\", ToString(byte(65))= " + got)
		}
	})
	t.Run("rune", func(t *testing.T) {
		if got := ToString(rune(65)); got != "A" {
			t.Errorf("ToString(rune(65)) != \"A\" ToString(rune(65)) = " + got)
		}
	})
}

func TestToASCIICode(t *testing.T) {
	t.Run("string_a", func(t *testing.T) {
		if ToASCIICode("a") != 97 {
			t.Errorf("ToASCIICode(\"a\") != 97")
		}
	})
	t.Run("string_b", func(t *testing.T) {
		if ToASCIICode("b") != 98 {
			t.Errorf("ToASCIICode(\"b\") != 98")
		}
	})
	t.Run("string_c", func(t *testing.T) {
		if ToASCIICode("c") != 99 {
			t.Errorf("ToASCIICode(\"c\") != 99")
		}
	})
	t.Run("string_Z", func(t *testing.T) {
		if ToASCIICode("Z") != 90 {
			t.Errorf("ToASCIICode(\"Z\") != 90")
		}
	})
	t.Run("byte", func(t *testing.T) {
		if ToASCIICode(byte(1)) != 1 {
			t.Errorf("ToASCIICode(byte(1)) != 1")
		}
	})
	t.Run("rune", func(t *testing.T) {
		if ToASCIICode(rune(1)) != 1 {
			t.Errorf("ToASCIICode(rune(1)) != 1")
		}
	})
}

func TestASCIIIntToChar(t *testing.T) {
	t.Run("string_a", func(t *testing.T) {
		if ASCIIIntToChar(97) != "a" {
			t.Errorf("ASCIIIntToChar(97) != \"a\"")
		}
	})
	t.Run("string_b", func(t *testing.T) {
		if ASCIIIntToChar(98) != "b" {
			t.Errorf("ASCIIIntToChar(98) != \"b\"")
		}
	})
	t.Run("string_c", func(t *testing.T) {
		if ASCIIIntToChar(99) != "c" {
			t.Errorf("ASCIIIntToChar(99) != \"c\"")
		}
	})
	t.Run("string_Z", func(t *testing.T) {
		if ASCIIIntToChar(90) != "Z" {
			t.Errorf("ASCIIIntToChar(90) != \"Z\"")
		}
	})
}
