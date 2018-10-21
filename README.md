# Base Converter (radix)

Base converter for various numeral systems plus custom alphabets.

## Supported numeral systems

 * Binary
 * Quaternary
 * Octal
 * Hexadecimal
 * Base62
 * Your own custom alphabet

## Example

```go
    // Binary
    encoded := Encode(123, Binary)
    decoded := Decode("1111011", Binary)

    // Base64
    encoded := Encode(18446744073709551615, Base62)
    decoded := Decode("LygHa16AHYF", Base62)

    // Custom alphabet
    alphabet := []string{"+", "_", "-", "(", ")", "[", "]"}
    encoded := Encode(18446744073709551615, alphabet))
    decoded := Decode(")[+_-+-_[--[-(_()_()]+_", alphabet)
```
