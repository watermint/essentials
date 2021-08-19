package ehex

import "essentials/eidiom"

func Parse(hex string) ([]byte, error) {
	if len(hex)%2 == 1 {
		return nil, eidiom.ErrorParseInvalidFormat
	}
	hr := []rune(hex)
	s := len(hr) / 2
	d := make([]byte, s)

	for i := 0; i < s; i++ {
		hi := parseSingleHex(hr[i*2])
		lo := parseSingleHex(hr[i*2+1])
		if hi > 0x10 || lo > 0x10 {
			return nil, eidiom.ErrorParseInvalidFormat
		}
		d[i] = hi<<4 | lo
	}
	return d, nil
}

func parseSingleHex(hex rune) byte {
	switch hex {
	case '0', '０':
		return 0x00
	case '1', '１':
		return 0x01
	case '2', '２':
		return 0x02
	case '3', '３':
		return 0x03
	case '4', '４':
		return 0x04
	case '5', '５':
		return 0x05
	case '6', '６':
		return 0x06
	case '7', '７':
		return 0x07
	case '8', '８':
		return 0x08
	case '9', '９':
		return 0x09
	case 'a', 'A', 'ａ', 'Ａ':
		return 0x0a
	case 'b', 'B', 'ｂ', 'Ｂ':
		return 0x0b
	case 'c', 'C', 'ｃ', 'Ｃ':
		return 0x0c
	case 'd', 'D', 'ｄ', 'Ｄ':
		return 0x0d
	case 'e', 'E', 'ｅ', 'Ｅ':
		return 0x0e
	case 'f', 'F', 'ｆ', 'Ｆ':
		return 0x0f
	default:
		return 0xff
	}
}
