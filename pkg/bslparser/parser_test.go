package bslparser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_real(t *testing.T) {

	path := "C:/Users/a.sosnoviy.VOSHOD/go/dev/go1cover/testresourse/cf/designer"
	p := New()
	fillse, _ := p.Files(path)
	p.Parse(fillse)

	for k, v := range p.LinesToCover {
		fmt.Println(k, v)
	}

}

func Test_nofile(t *testing.T) {

	defer func() {
        if r := recover(); r != nil {
           t.Fail()
        }
    }()

	p := New()
	p.Parsefile("nofilebsl")

}

func Test_Test(t *testing.T) {

	a := New().Parsefile("./tests/test.bsl")
	r := []int{13, 17, 19, 22, 23, 24, 27, 29, 30, 31, 33, 34, 35, 38, 39, 41, 42, 45, 46, 50, 51, 52, 54,
		56, 58, 60, 66, 67, 69, 70, 71, 73, 74, 75, 76, 78, 79, 80, 81, 82, 84, 92, 93, 94, 98, 103, 105,
		112}

	assert.EqualValues(t, r, a)

}

func Test_linesToCover(t *testing.T) {

	a := New().Parsefile("./tests/tocover.bsl")
	r := []int{13, 17, 19, 22, 23, 24, 27, 29, 30, 31, 33, 34, 35, 38, 39, 41, 42, 45, 46, 50, 51, 52, 54,
		56, 58, 60, 66, 67, 69, 70, 71, 73, 74, 75, 76, 78, 79, 80, 81, 82, 84, 92, 93, 94, 98, 103, 105,
		112}
	assert.EqualValues(t, r, a)

}

func Test_simple(t *testing.T) {

	a := New().Parsefile("./tests/simple.bsl")
	r := []int{3, 8, 15, 18, 19, 20, 21, 23, 24, 29, 30, 31, 32, 33, 34, 35, 36, 37, 41, 42, 48,
		50, 53, 57, 60, 63, 66, 67, 68, 69, 70, 72}
	assert.EqualValues(t, r, a)

}

func Test_if(t *testing.T) {

	a := New().Parsefile("./tests/if.bsl")
	r := []int{3, 8, 10, 16, 17, 19, 21, 22, 23}
	assert.EqualValues(t, r, a)

}

func Test_assigment(t *testing.T) {

	a := New().Parsefile("./tests/assigment.bsl")
	r := []int{5, 11, 18, 22, 26, 32, 38, 44, 46, 52, 58, 63, 68, 70, 79, 84, 92, 102, 106, 110, 114,
		117, 121, 125, 127, 128}
	assert.EqualValues(t, r, a)

}

func Test_do(t *testing.T) {

	a := New().Parsefile("./tests/do.bsl")
	r := []int{4, 5, 6, 7, 8, 12, 16, 22, 23, 25, 30, 31, 32, 33, 34, 35, 36, 37, 39, 43, 44, 53, 54,
		56, 61, 64, 65, 66, 67, 68, 70, 75, 76, 77, 78, 79, 80, 82, 85, 86, 89, 95, 96, 97, 99, 102, 103, 106,
		107, 109, 110, 111, 112}
	assert.EqualValues(t, r, a)

}

func Test_other(t *testing.T) {

	a := New().Parsefile("./tests/other.bsl")
	r := []int{4, 10, 15, 18, 19, 21, 34, 48, 50, 53, 55, 59, 61, 62, 64, 65,
		70, 72, 76, 80, 84, 87, 91, 92, 94, 95, 96}
	assert.EqualValues(t, r, a)

}

func Test_notcovered(t *testing.T) {

	a := New().Parsefile("./tests/notcovered.bsl")
	r := []int{6, 7,
		8, // Лишняя
		9, 13, 14, 16,
		17, // Лишняя
		19, 21, 22, 23}
	assert.EqualValues(t, r, a)

}

func Test_filtered(t *testing.T) {

	a := New().Parsefile("./tests/filtered.bsl")
	r := []int{3, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 23, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 62, 63, 65, 70, 73, 75, 77, 79, 81, 83, 85, 87, 89, 91, 93, 95, 97, 99, 101, 103, 105, 107, 109, 111, 113, 115, 117, 119, 122, 126, 128, 130, 132, 134, 135, 136, 137, 139, 143, 146, 149, 152, 155, 156, 157, 158}
	assert.EqualValues(t, r, a)

}

func Test_opcode(t *testing.T) {

	a := New().Parsefile("./tests/opcode.bsl")

	r := []int{3, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 111, 113, 115, 117, 119, 121, 123, 125, 127, 129, 131, 133, 135, 137, 139, 141, 143, 145, 147, 149, 152, 153, 155, 156, 159, 161}
	assert.EqualValues(t, r, a)

}

func Test_error(t *testing.T) {

	a := New().Parsefile("./tests/error.bsl")
	r := []int{4,
		//	6, // Почемуто нету
		8,
	}
	assert.EqualValues(t, r, a)

}
