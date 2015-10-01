# go-TEMPer
read TEMPer USB HID devices (USB ID 0c45:7401) via gousb

 * based on pcsensor.c by Juan Carlos Perez (c) 2011 (cray@isp-sl.com)
 * based on Temper.c by Robert Kavaler (c) 2009 (relavak.com)

This is a simple go program that reads the temperature from TEMPer USB
thermometers using the gousb libusb-1.0 wrapper. It might be an easy to
follow example of gousb in practice.

sudo go run Temper.go

will print out a timestamp and the temperature.

gousb does not include the iProduct field, so this version assumes
a TEMPer1F_V1.3 with offset 4 for the USB read buffer.

