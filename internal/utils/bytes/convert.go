package bytes

import "unsafe"

func ToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
