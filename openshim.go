// Copyright 2018 Esote. All rights reserved. Use of this source code is
// governed by an MIT license that can be found in the LICENSE file.

package openshim

//	#ifdef __OpenBSD__
//	#include <unistd.h>
//	#endif
//
//	int
//	shim_pledge(const char *promises, const char *execpromises) {
//	#ifdef __OpenBSD__
//		return pledge(promises, execpromises);
//	#else
//		return -1;
//	#endif
//	}
//
//	int
//	shim_unveil(const char *path, const char *permissions) {
//	#ifdef __OpenBSD__
//		return unveil(path, permissions);
//	#else
//		return -1;
//	#endif
//	}
import "C"

// Deprecated: Pledge syscall, error is errno value
func Pledge(promises, execpromises string) (int, error) {
	i, err := C.shim_pledge(C.CString(promises), C.CString(execpromises))
	return int(i), err
}

// Deprecated: Unveil syscall, error is errno value
func Unveil(path, permissions string) (int, error) {
	i, err := C.shim_unveil(C.CString(path), C.CString(permissions))
	return int(i), err
}
