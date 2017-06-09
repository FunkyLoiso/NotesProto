package core

import (
	// "bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestSpaceCompressor(t *testing.T) {
	tests := []struct{ text, exp string }{
		{"Безопасноть и забота", "Безопасноть и забота"},
		{"1 22  333", "1 22 333"},
		{" 1 2 3 ", " 1 2 3 "},
		{"  島  民  ", " 島 民 "},
		{"          ", " "},
		{"1  3  5  7  9  1  3  5  7  9  1  3  5  7  9  1  3  5", "1 3 5 7 9 1 3 5 7 9 1 3 5 7 9 1 3 5"},
		{"1908年（明治41年）にマックス・プランクが、相対性理論という語を作り、どのように特殊相対性理論（のちに一般相対性理論も）相対性原理を適用するのかを説明した。", "1908年（明治41年）にマックス・プランクが、相対性理論という語を作り、どのように特殊相対性理論（のちに一般相対性理論も）相対性原理を適用するのかを説明した。"},
		{"", ""},
		{"\n\n\n", "\n"},
		{"1\n2  \n 3\n  4 \n", "1\n2 3\n4 "},
	}

	for _, v := range tests {
		cp := NewSpaceCompressor(strings.NewReader(v.text))
		data, err := ioutil.ReadAll(cp)
		if err != nil {
			t.Error(err.Error())
		}
		if result := string(data); result != v.exp {
			t.Errorf("SpaceCompressor failed for '%v'\nexpected: '%v'\nresult:   '%v'", v.text, v.exp, result)
		}
	}
}

func TestMakeTitle(t *testing.T) {
	tests := []struct{ text, title string }{
		{"Собака на работе.", "Собака на работе."},
		{`Кошка
		    на
		    природе`, "Кошка"},
		{`
		    islander`, ""},
		{"島民", "島民"},
		{"12345678901234567890123456789012345", "123456789012345678901234567890"},
		{"12 345678901234567890123456789012345", "12"},
		{"ab cdefghijklmnopqrstuvwxyzABC", "ab cdefghijklmnopqrstuvwxyzABC"},
		{"ab cdefghijklmnopqrstuvwxyzABCD", "ab"},
		{"1  3  5  7  9  1  3  5  7  9  1  3  5  7  9  1  3  5", "1 3 5 7 9 1 3 5 7 9 1 3 5 7 9"},
		{"1908年（明治41年）にマックス・プランクが、相対性理論という語を作り、どのように特殊相対性理論（のちに一般相対性理論も）相対性原理を適用するのかを説明した。", "1908年（明治41年）にマックス・プランクが、"},
	}

	for _, v := range tests {
		result := MakeTitle(v.text)
		if result != v.title {
			t.Errorf("MakeTitle failed for '%v'\nexpected: '%v'\nresult:   '%v'", v.text, v.title, result)
		}
	}
}
