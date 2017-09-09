package bits

import (
    "testing"
    "bytes"
    "fmt"
)

/*
unsigned integer = 9
binary = 1001 (4 bits)
exp-golomb code = 000 1 010 (7 bits)

unsigned integer = 50
binary = 110010 (6 bits)
exp-golomb code = 00000 1 10011 (11 bits)
*/
var testData = []struct{
    ExpGolombBit []byte
    BitNumber int
    UnsignedInt uint
}{
    {[]byte{0x0a}, 7, 9},
    {[]byte{0x00, 0x33}, 11, 50},
}

func TestGolombBitReader_ReadBit(t *testing.T) {

    for _, data := range testData {
        r := &GolombBitReader{R: bytes.NewReader(data.ExpGolombBit)}
        for i := 0; i < len(data.ExpGolombBit) * 8 - data.BitNumber; i++ {
            b, _ := r.ReadBit()
            fmt.Println("drop bit", b)
        }
        res, _ := r.ReadExponentialGolombCode()
        if res != data.UnsignedInt {
            t.Fatalf("read golomb code failed, exp=%v, actual=%v", data.UnsignedInt, res)
        }
    }
}
