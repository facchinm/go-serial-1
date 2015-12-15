package main

import (
	".."
	"fmt"
)

func main() {
	ports, err := serial.ListPorts()
	if err != nil {
		return
	}

	for _, info := range ports {
		if vid, pid, err := info.USBVIDPID(); err == nil {
			fmt.Printf("%s | %04X:%04X | %s\n", info.Name(), vid, pid, info.USBSerialNumber())
		}
	}

	return
}
