package taggedptr

import (
	"testing"
	"unsafe"
)

type TestStruct struct {
	Field int
}

func TestTag(t *testing.T) {
	s := &TestStruct{42}

	initialPtr := unsafe.Pointer(s)
	newPtr, _ := Tag(unsafe.Pointer(s), 3)
	s = (*TestStruct)(newPtr)

	if unsafe.Pointer(s) != unsafe.Pointer(uintptr(initialPtr)|3) {
		t.Fatal("wrong tag")
	}

	if _, err := Tag(unsafe.Pointer(s), 10); err == nil {
		t.Fatal("expected error")
	}
}

func TestGetPointer(t *testing.T) {
	s := &TestStruct{42}
	initialPtr := unsafe.Pointer(s)
	newPtr, _ := Tag(unsafe.Pointer(s), 3)
	s = (*TestStruct)(newPtr)

	if ptr := GetPointer(unsafe.Pointer(s)); ptr != initialPtr {
		t.Fatalf("got wrong pointer. expected: %p, got: %p", initialPtr, ptr)
	}
}

func TestGetTag(t *testing.T) {
	s := &TestStruct{42}
	newPtr, _ := Tag(unsafe.Pointer(s), 3)
	s = (*TestStruct)(newPtr)

	if tag := GetTag(unsafe.Pointer(s)); tag != 3 {
		t.Fatalf("got wrong pointer. expected: %d, got: %d", 3, tag)
	}
}

func TestCompareAndSwap(t *testing.T) {
	s := &TestStruct{42}
	casPtr := unsafe.Pointer(s)
	initialPtr := unsafe.Pointer(s)

	if _ = CompareAndSwap(&casPtr, casPtr, casPtr, 0, 1); casPtr != unsafe.Pointer(uintptr(initialPtr)|1) {
		t.Fatal("wrong cas")
	}
}

func TestGet(t *testing.T) {
	s := &TestStruct{42}
	initialPtr := unsafe.Pointer(s)
	newPtr, _ := Tag(unsafe.Pointer(s), 3)
	s = (*TestStruct)(newPtr)

	if ptr, tag := Get(unsafe.Pointer(s)); ptr != initialPtr || tag != 3 {
		t.Fatalf("wrong get. expected: %p %d, got: %p %d", initialPtr, 3, ptr, tag)
	}
}
