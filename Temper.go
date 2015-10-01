package main
/*
 * Temper.go by Thomas Burke (c) 2015 (tburke@tb99.com)
 * based on pcsensor.c by Juan Carlos Perez (c) 2011 (cray@isp-sl.com)
 * based on Temper.c by Robert Kavaler (c) 2009 (relavak.com)
*/

import (
	"github.com/kylelemons/gousb/usb"
	"log"
	"fmt"
)

func temperature() (float64, error) {
	ctx := usb.NewContext()
	defer ctx.Close()

	devs, err := ctx.ListDevices(func(desc *usb.Descriptor) bool {
		return desc.Vendor == 0x0c45 && desc.Product == 0x7401
	})

	defer func() {
		for _, d := range devs {
			d.Close()
		}
	}()

	if err != nil {
		return 0.0, err
	}

	if len(devs) == 0 {
		return 0.0, fmt.Errorf("No thermometers found.")
	}

	dev := devs[0]
	if err = dev.SetConfig(1); err != nil {
		return 0.0, err
	}

	ep, err := dev.OpenEndpoint(1, 1, 0, 0x82)
	if err != nil {
		return 0.0, err
	}
	if _, err = dev.Control(0x21, 0x09, 0x0200, 0x01, []byte{0x01, 0x80, 0x33, 0x01, 0x00, 0x00, 0x00, 0x00}); err != nil {
		return 0.0, err
	}
	buf := make([]byte, 8)
	if _, err = ep.Read(buf); err != nil {
		return 0.0, err
	}
	return float64(buf[4]) + float64(buf[5])/256, nil
}

func main() {
	c, err := temperature()
	if err == nil {
		log.Printf("Temperature: %.2fF %.2fC\n", 9.0/5.0*c+32, c)
	} else {
		log.Fatalf("Failed: %s", err)
	}
}
