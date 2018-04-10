package namecase

import "unsafe"

func toString(d []byte) string {
	return *(*string)(unsafe.Pointer(&d))
}
