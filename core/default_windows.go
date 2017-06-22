// +build windows

package core

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

func getDefaultEditor() string {
	sh32, err := syscall.LoadLibrary("Shell32.dll")
	if err != nil {
		log.Println("LoadLibrary for 'Shell32.dll' failed", err)
		return ""
	}
	findExecutable, err := syscall.GetProcAddress(sh32, "FindExecutableW")
	if err != nil {
		log.Println("GetProcAddress for 'FindExecutableW' failed", err)
		return ""
	}

	filePath := os.TempDir() + "\\" + strconv.Itoa(rand.Int()) + "_.txt"
	err = ioutil.WriteFile(filePath, []byte("dummy text to get default text editor"), 0600)
	if err != nil {
		log.Printf("Failed to create dummy text file '%v': %v", filePath, err)
		return ""
	}
	defer os.Remove(filePath)

	namePtr, err := syscall.UTF16PtrFromString(filePath)
	if err != nil {
		log.Printf("UTF16PtrFromString failed for '%v': %v", filePath, err)
		return ""
	}
	outBuff := make([]uint16, syscall.MAX_PATH)
	ret, _, err := syscall.Syscall(
		findExecutable, 3,
		uintptr(unsafe.Pointer(namePtr)),
		uintptr(0),
		uintptr(unsafe.Pointer(&outBuff[0])))

	// err is ignored because for some reason it is always set to
	// 'An attempt was made to reference a token that does not exist.', while the returned value seems to be fine
	if err != nil {
		log.Println("Syscall call for 'FindExecutableW' returned error, ignoring:", err)
	}

	if ret <= 32 {
		log.Printf("Syscall for 'FindExecutableW' failed with code %v", ret)
		return ""
	}

	for i, v := range outBuff {
		if v == 0 {
			outBuff = outBuff[0:i]
			break
		}
	}
	return string(utf16.Decode(outBuff))
}
