// This module contains methods to tag the pointer.
// Tagged pointer is a pointer with additional data associate with it.
// It's possible because data must be word aligned,
// hence least significant bits could be used to store some data.
// It's useful in lock-free programming,
// to store some data in pointer atomically using CAS instructions.
package taggedptr

import (
	"errors"
	"sync/atomic"
	"unsafe"
)

const MAX_TAG_SIZE = 3

// Return tagged pointer.
func Tag(ptr unsafe.Pointer, tag uint) (unsafe.Pointer, error) {
	if tag > MAX_TAG_SIZE {
		return nil, errors.New("Too large tag")
	}

	ptr = unsafe.Pointer(uintptr(ptr) | uintptr(tag))

	return ptr, nil
}

// Get the current value of the pointer.
func GetPointer(ptr unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr(ptr) &^ uintptr(MAX_TAG_SIZE))
}

// Return the current value of the tag.
func GetTag(ptr unsafe.Pointer) uint {
	return uint(uintptr(ptr) & uintptr(MAX_TAG_SIZE))
}

// Get pair of pointer and tag values.
func Get(ptr unsafe.Pointer) (unsafe.Pointer, uint) {
	return GetPointer(ptr), GetTag(ptr)
}

// Compare and swap tagged pointer.
func CompareAndSwap(addr *unsafe.Pointer, oldPtr, newPtr unsafe.Pointer, oldTag, newTag uint) bool {
	var err error

	if oldPtr, err = Tag(oldPtr, oldTag); err != nil {
		return false
	}

	if newPtr, err = Tag(newPtr, newTag); err != nil {
		return false
	}

	return atomic.CompareAndSwapPointer(addr, oldPtr, newPtr)
}
