package dequeue

import (
	"testing"
)

//func init() {
//	log.SetFlags(log.Lshortfile)
//	var logFileName = "test.log"
//	var logFile, err = os.Create(logFileName)
//	if err != nil {
//		log.Fatal(errors.Wrapf(err, "failed to os.Create(%q)", logFileName))
//	}
//	log.SetOutput(logFile)
//}

func TestCreateDequeue(t *testing.T) {
	var d = New()
	if d.Next() != d.Prev() {
		t.Fatal("d.Next() != d.Prev()")
	}
	if d.Next() != &d.Node {
		t.Fatal("d.Next() != &d.Node")
	}
}

func TestPush(t *testing.T) {
	var d = New()
	if d.Push("a") != d {
		t.Fatal("d.Push(\"a\") != d")
	}
	if d.Next() != d.Prev() {
		t.Fatal("d.Next() != d.Prev()")
	}
	if d.Next().Data() != "a" {
		t.Fatal("d.Next().Data() != \"a\"")
	}
}

func TestPop(t *testing.T) {
	var d = New()

	if d.Pop() != nil {
		t.Fatal("d.Pop() != nil")
	}

	d.Push("a").Push("b").Push("c")
	if d.Pop() != "c" {
		t.Fatal("d.Pop() != \"c\"")
	}

	if d.Next().Data() != "a" {
		t.Fatal("d.Next().Data() != \"a\"")
	}
	if d.Prev().Data() != "b" {
		t.Fatal("d.Prev().Data() != \"b\"")
	}
}

func TestUnshift(t *testing.T) {
	var d = New()

	if d.Unshift("a") != d {
		t.Fatal("d.Unshift(\"a\") != d")
	}

	if d.Next() != d.Prev() {
		t.Fatal("d.Next() != d.Prev()")
	}
	if d.Prev().Data() != "a" {
		t.Fatal("d.Prev().Data() != \"a\"")
	}
}

func TestShift(t *testing.T) {
	var d = New()

	if d.Shift() != nil {
		t.Fatal("d.Shift() != nil")
	}

	d.Unshift("c").Unshift("b").Unshift("a")
	if d.Shift() != "a" {
		t.Fatal("d.Shift() != \"a\"")
	}

	if d.Next().Data() != "b" {
		t.Fatal("d.Next().Data() != \"b\"")
	}
	if d.Prev().Data() != "c" {
		t.Fatal("d.Prev().Data() != \"a\"")
	}
}

func TestRange(t *testing.T) {
	var d = New().Push("a").Push("b").Push("c")

	var str string
	d.Range(func(data interface{}) bool {
		var s = data.(string)
		str = str + s
		return true
	})
	//fmt.Printf("str = %q\n", str) //needs `go test -v` to see this output
	if str != "abc" {
		t.Fatalf("str,%q != \"abc\"", str)
	}

	str = ""
	d.Range(func(data interface{}) bool {
		var s = data.(string)
		str = str + s
		if s == "b" {
			return false
		}
		return true
	})
	//fmt.Printf("str = %q\n", str) //needs `go test -v` to see this output
	if str != "ab" {
		t.Fatalf("str,%q != \"ab\"", str)
	}
}

func TestLen(t *testing.T) {
	var d = New().Push("a").Push("b").Push("c")

	if d.Len() != 3 {
		t.Fatal("d.Len() != 3")
	}

	var v0 = d.Shift()
	if v0 != "a" {
		t.Fatal("v0 != \"a\"")
	}

	var v1 = d.Pop()
	if v1 != "c" {
		t.Fatal("v1 != \"c\"")
	}

	if d.Len() != 1 {
		t.Fatal("d.Len() != 1")
	}
}

func TestEmpty(t *testing.T) {
	var d = New().Push("a").Push("b").Push("c")

	if d.Len() != 3 {
		t.Fatal("d.Len() != 3")
	}

	var v0 = d.Shift()
	if v0 != "a" {
		t.Fatal("v0 != \"a\"")
	}

	var v1 = d.Pop()
	if v1 != "c" {
		t.Fatal("v1 != \"c\"")
	}

	if d.Len() != 1 {
		t.Fatal("d.Len() != 1")
	}

	d.Pop()
	if !d.Empty() {
		t.Fatal("!d.Empty()")
	}
}

func TestFirst(t *testing.T) {
	var d = New().Push("a").Push("b").Push("c")

	var v0 = d.First().Data()
	if v0 != "a" {
		t.Fatal("v0 != \"a\"")
	}

	d.Shift()
	var v1 = d.First().Data()
	if v1 != "b" {
		t.Fatal("v1 != \"b\"")
	}

	d.Pop()
	var v2 = d.First().Data()
	if v2 != "b" {
		t.Fatal("v2 != \"b\"")
	}

	d.Shift()
	if d.First() != nil {
		t.Fatal("d.First() != nil")
	}
}

func TestLast(t *testing.T) {
	var d = New().Push("a").Push("b").Push("c")

	var v0 = d.Last().Data()
	if v0 != "c" {
		t.Fatal("v0 != \"c\"")
	}

	d.Shift()
	var v1 = d.Last().Data()
	if v1 != "c" {
		t.Fatal("v1 != \"c\"")
	}

	d.Pop()
	var v2 = d.Last().Data()
	if v2 != "b" {
		t.Fatal("v2 != \"b\"")
	}

	d.Shift()
	if d.Last() != nil {
		t.Fatal("d.Last() != nil")
	}
}
