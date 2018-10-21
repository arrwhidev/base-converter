package radix

import (
	"testing"

	utils "github.com/arrwhidev/testutils-golang"
)

func TestBinaryEncodeAndDecode(t *testing.T) {
	input := uint64(18446744073709551615)
	decoded := Decode(Encode(input, Binary), Binary)
	utils.AssertEqual(t, input, decoded)
}

func TestBinaryEncoding(t *testing.T) {
	utils.AssertEqual(t, "0", Encode(0, Binary))
	utils.AssertEqual(t, "1111011", Encode(123, Binary))
	utils.AssertEqual(t, "100110001001011001111111", Encode(9999999, Binary))
	utils.AssertEqual(t, "1111111111111111111111111111111111111111111111111111111111111111", Encode(18446744073709551615, Binary))
}

func TestQuaternaryEncoding(t *testing.T) {
	utils.AssertEqual(t, "0", Encode(0, Quaternary))
	utils.AssertEqual(t, "1323", Encode(123, Quaternary))
	utils.AssertEqual(t, "212021121333", Encode(9999999, Quaternary))
	utils.AssertEqual(t, "33333333333333333333333333333333", Encode(18446744073709551615, Quaternary))
}

func TestOctalEncoding(t *testing.T) {
	utils.AssertEqual(t, "0", Encode(0, Octal))
	utils.AssertEqual(t, "173", Encode(123, Octal))
	utils.AssertEqual(t, "46113177", Encode(9999999, Octal))
	utils.AssertEqual(t, "1777777777777777777777", Encode(18446744073709551615, Octal))
}

func TestHexadecimalEncoding(t *testing.T) {
	utils.AssertEqual(t, "0", Encode(0, Hexadecimal))
	utils.AssertEqual(t, "7B", Encode(123, Hexadecimal))
	utils.AssertEqual(t, "98967F", Encode(9999999, Hexadecimal))
	utils.AssertEqual(t, "FFFFFFFFFFFFFFFF", Encode(18446744073709551615, Hexadecimal))
}

func TestBase62Encoding(t *testing.T) {
	utils.AssertEqual(t, "0", Encode(0, Base62))
	utils.AssertEqual(t, "1z", Encode(123, Base62))
	utils.AssertEqual(t, "fxSJ", Encode(9999999, Base62))
	utils.AssertEqual(t, "LygHa16AHYF", Encode(18446744073709551615, Base62))
}

func TestCustomAlphabetEncoding(t *testing.T) {
	customAlphabet := []string{"+", "_", "-", "(", ")", "[", "]"}
	utils.AssertEqual(t, "+", Encode(0, customAlphabet))
	utils.AssertEqual(t, "-()", Encode(123, customAlphabet))
	utils.AssertEqual(t, "_[+]]]()-", Encode(9999999, customAlphabet))
	utils.AssertEqual(t, ")[+_-+-_[--[-(_()_()]+_", Encode(18446744073709551615, customAlphabet))
}

func TestBinaryDecoding(t *testing.T) {
	utils.AssertEqual(t, uint64(0), Decode("0", Binary))
	utils.AssertEqual(t, uint64(123), Decode("1111011", Binary))
	utils.AssertEqual(t, uint64(9999999), Decode("100110001001011001111111", Binary))
	utils.AssertEqual(t, uint64(18446744073709551615), Decode("1111111111111111111111111111111111111111111111111111111111111111", Binary))
}

func TestQuaternaryDecoding(t *testing.T) {
	utils.AssertEqual(t, uint64(0), Decode("0", Quaternary))
	utils.AssertEqual(t, uint64(123), Decode("1323", Quaternary))
	utils.AssertEqual(t, uint64(9999999), Decode("212021121333", Quaternary))
	utils.AssertEqual(t, uint64(18446744073709551615), Decode("33333333333333333333333333333333", Quaternary))
}

func TestOctalDecoding(t *testing.T) {
	utils.AssertEqual(t, uint64(0), Decode("0", Octal))
	utils.AssertEqual(t, uint64(123), Decode("173", Octal))
	utils.AssertEqual(t, uint64(9999999), Decode("46113177", Octal))
	utils.AssertEqual(t, uint64(18446744073709551615), Decode("1777777777777777777777", Octal))
}

func TestHexadecimalDecoding(t *testing.T) {
	utils.AssertEqual(t, uint64(0), Decode("0", Hexadecimal))
	utils.AssertEqual(t, uint64(123), Decode("7B", Hexadecimal))
	utils.AssertEqual(t, uint64(9999999), Decode("98967F", Hexadecimal))
	utils.AssertEqual(t, uint64(18446744073709551615), Decode("FFFFFFFFFFFFFFFF", Hexadecimal))
}

func TestBase62Decoding(t *testing.T) {
	utils.AssertEqual(t, uint64(0), Decode("0", Base62))
	utils.AssertEqual(t, uint64(123), Decode("1z", Base62))
	utils.AssertEqual(t, uint64(9999999), Decode("fxSJ", Base62))
	utils.AssertEqual(t, uint64(18446744073709551615), Decode("LygHa16AHYF", Base62))
}

func TestCustomAlphabetDecoding(t *testing.T) {
	customAlphabet := []string{"+", "_", "-", "(", ")", "[", "]"}
	utils.AssertEqual(t, uint64(0), Decode("+", customAlphabet))
	utils.AssertEqual(t, uint64(123), Decode("-()", customAlphabet))
	utils.AssertEqual(t, uint64(9999999), Decode("_[+]]]()-", customAlphabet))
	utils.AssertEqual(t, uint64(18446744073709551615), Decode(")[+_-+-_[--[-(_()_()]+_", customAlphabet))
}
