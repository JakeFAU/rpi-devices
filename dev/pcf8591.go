/*
Package dev ...

PCF8591 is the driver of PCF8591 module.

connect to raspberry pi:
- VCC: pin 1 or any 3.3v pin
- GND: pin 9 or and GND pin
- SDA: pin 3 (SDA)
- SCL: pin 5 (SCL)

Jumper:
- remove jumpers on P4 & P5, keep the jumper on P6

Config Your Pi:
1. $ sudo apt-get install -y python-smbus
2. $ sudo apt-get install -y i2c-tools
3. $ sudo raspi-config
4. 	-> [5 interface options] -> [p5 i2c] ->[yes] -> [ok]
5. $ sudo reboot now
6. check: $ sudo i2cdetect -y 1
	it works if you saw following message:
	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	     0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f
	00:          -- -- -- -- -- -- -- -- -- -- -- -- --
	10: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	20: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	30: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	40: -- -- -- -- -- -- -- -- 48 -- -- -- -- -- -- --
	50: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	60: -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
	70: -- -- -- -- -- -- -- --
	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*/
package dev

import (
	"log"

	"golang.org/x/exp/io/i2c"
)

const (
	pcf8591DevFile = "/dev/i2c-1"
	addrPCF8591    = 0x48
	ctrAIN0        = 0x40
	ctrAIN1        = 0x41
	ctrAIN2        = 0x42
	ctrAIN3        = 0x43
)

// PCF8591 ...
type PCF8591 struct {
	dev *i2c.Device
}

// NewPCF8591 ...
func NewPCF8591() (*PCF8591, error) {
	dev, err := i2c.Open(&i2c.Devfs{Dev: pcf8591DevFile}, addrPCF8591)
	if err != nil {
		return nil, err
	}
	return &PCF8591{
		dev: dev,
	}, nil
}

// ReadAIN0 ...
func (m *PCF8591) ReadAIN0() []byte {
	if err := m.dev.Write([]byte{ctrAIN0}); err != nil {
		log.Printf("write AIN0 error: %v", err)
		return []byte{}
	}
	data := make([]byte, 1)
	if err := m.dev.Read(data); err != nil {
		log.Printf("read AIN0 error: %v", err)
		return []byte{}
	}
	log.Printf("ain0, len: %v, data: %v", len(data), data)
	return data
}

// ReadAIN1 ...
func (m *PCF8591) ReadAIN1() []byte {
	if err := m.dev.Write([]byte{ctrAIN1}); err != nil {
		log.Printf("write AIN1 error: %v", err)
		return []byte{}
	}
	data := make([]byte, 1)
	if err := m.dev.Read(data); err != nil {
		log.Printf("read AIN1 error: %v", err)
		return []byte{}
	}
	log.Printf("ain1, len: %v, data: %v", len(data), data)
	return data
}

// ReadAIN2 ...
func (m *PCF8591) ReadAIN2() []byte {
	if err := m.dev.Write([]byte{ctrAIN2}); err != nil {
		log.Printf("write AIN2 error: %v", err)
		return []byte{}
	}
	data := make([]byte, 1)
	if err := m.dev.Read(data); err != nil {
		log.Printf("read AIN2 error: %v", err)
		return []byte{}
	}
	log.Printf("ain2, len: %v, data: %v", len(data), data)
	return data
}

// ReadAIN3 ...
func (m *PCF8591) ReadAIN3() []byte {
	if err := m.dev.Write([]byte{ctrAIN3}); err != nil {
		log.Printf("write AIN3 error: %v", err)
		return []byte{}
	}
	data := make([]byte, 1)
	if err := m.dev.Read(data); err != nil {
		log.Printf("read AIN3 error: %v", err)
		return []byte{}
	}
	return data
}

// Close ...
func (m *PCF8591) Close() {
	m.dev.Close()
}
