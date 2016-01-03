[![Build Status](https://travis-ci.org/alexyer/taggedptr.svg)](https://travis-ci.org/alexyer/taggedptr)
[![Coverage Status](https://coveralls.io/repos/alexyer/taggedptr/badge.svg?branch=master&service=github)](https://coveralls.io/github/alexyer/taggedptr?branch=master)

# Description
This module contains methods to tag the pointer.
Tagged pointer is a pointer with additional data associate with it.
It's possible because data must be word aligned,
hence least significant bits could be used to store some data.
It's useful in lock-free programming,
to store some data in pointer atomically using CAS instructions.

## Examples
```go
type TestStruct struct {
	Field int
}

s := &TestStruct{42}

// Tag pointer and assign new value
newPtr, _ := Tag(unsafe.Pointer(s), 3)
s = (*TestStruct)(newPtr)

// Get clear pointer
initialPtr := GetPointer(unsafe.Pointer(s))

// Get tag
tag := GetTag(unsafe.Pointer(s))

// Tag, compare and swap pointer
casPtr := unsafe.Pointer(s)
CompareAndSwap(&casPtr, casPtr, casPtr, 0, 1)
```
