package safe

import (
	"errors"
	"testing"
)

type InnerStruct struct {
	InnerField string `json:"innerField" bson:"innerField"`
}

type TestingWithStruct struct {
	OuterField  string `bson:"outerField" json:"outerField"`
	InnerStruct `bson:"innerStruct" json:"innerStruct"`
}

func BenchmarkOkSlicesStruct(b *testing.B) {
	preallocate := []TestingWithStruct{
		{
			OuterField:  "croot",
			InnerStruct: InnerStruct{"croot"},
		},
		{
			OuterField:  "croot",
			InnerStruct: InnerStruct{"croot"},
		},
	}
	for i := 0; i < b.N; i++ {
		val := Ok(preallocate)
		val.IsOk()
	}
	b.ReportAllocs()
}
func BenchmarkOkSlicesString(b *testing.B) {
	strings := []string{"", "", ""}
	for i := 0; i < b.N; i++ {
		val := Ok(strings)
		val.IsOk()
	}
	b.ReportAllocs()
}

func BenchmarkOkSlicesInt(b *testing.B) {
	ints := []int{0, 0, 0}
	for i := 0; i < b.N; i++ {
		val := Ok(ints)
		val.IsOk()
	}
	b.ReportAllocs()
}
func BenchmarkOkMapIntString(b *testing.B) {
	ints := map[int]string{1: "testing 1", 2: "testing 2", 3: "testing 3"}
	for i := 0; i < b.N; i++ {
		val := Ok(ints)
		val.IsOk()
	}
	b.ReportAllocs()
}

func BenchmarkResult_Err(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Err[int]("some fancy error message")
	}
	b.ReportAllocs()
}

func BenchmarkResult_Ok(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Ok(TestingWithStruct{})
		Ok(TestingWithStruct{})
		Ok(TestingWithStruct{})
		Ok(TestingWithStruct{})
		Ok(TestingWithStruct{})
	}
	b.ReportAllocs()
}

func BenchmarkResultTestOk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result_test_ok()
	}
	b.ReportAllocs()
}

func BenchmarkAsResultEmptyErr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := AsResult[TestingWithStruct](emptyStructErr())
		val.IsErr()
	}
	b.ReportAllocs()
}
func BenchmarkAsResultEmptyNoErr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := AsResult[TestingWithStruct](emptyStruct())
		val.IsOk()
	}
	b.ReportAllocs()
}

func BenchmarkResult_OkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Ok(23)
		res.IsOk()
	}
	b.ReportAllocs()
}

func BenchmarkResult_EmptyInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Ok(0)
		res.IsOk()
	}
	b.ReportAllocs()
}

func BenchmarkResult_OkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Ok("hello world")
		res.IsOk()
	}
	b.ReportAllocs()
}

func BenchmarkResult_EmptyString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Ok("")
		res.IsOk()
	}
	b.ReportAllocs()
}

func BenchmarkOption_Some(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Some[TestingWithStruct](TestingWithStruct{})
		Some[TestingWithStruct](TestingWithStruct{})
		Some[TestingWithStruct](TestingWithStruct{})
	}
	b.ReportAllocs()
}

func BenchmarkOption_None(b *testing.B) {
	for i := 0; i < b.N; i++ {
		None[int]()
	}
	b.ReportAllocs()
}

func BenchmarkOption_IsNone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		val := None[TestingWithStruct]()
		val.IsNone()
	}
	b.ReportAllocs()
}

func BenchmarkOption_IsSome(b *testing.B) {
	withStruct := TestingWithStruct{}
	for i := 0; i < b.N; i++ {
		val := Some(withStruct)
		val.IsSome()
	}
	b.ReportAllocs()
}

func BenchmarkErrorCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := AsResult(TestingWithStruct{}, errors.New("some fancy error message"))
		err.IsErr()

	}
	b.ReportAllocs()
}

func emptyStructErr() (data TestingWithStruct, err error) {
	err = errors.New("some fancy error message")
	return
}

func emptyStruct() (data TestingWithStruct, err error) {
	data = TestingWithStruct{
		OuterField:  "testing",
		InnerStruct: InnerStruct{InnerField: "testing2"},
	}
	return
}
