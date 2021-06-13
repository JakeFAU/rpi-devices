package main

import (
	"log"
	"time"

	"github.com/jakefau/rpi-devices/dev"
	"github.com/jakefau/rpi-devices/iot"
	"github.com/jakefau/rpi-devices/util"
)

func main() {
	gps := dev.NewGPS("/dev/ttyAMA0", 9600)
	// gps := dev.NewMockGPS("./dev/test/gps.csv")
	if gps == nil {
		log.Printf("[gpstracker]failed to new a gps device")
		return
	}
	logger := util.NewGPSLogger()
	if logger == nil {
		log.Printf("[gpstracker]failed to new a tracker")
		return
	}
	oneNetCfg := &iot.OneNetConfig{
		Token: iot.OneNetToken,
		API:   iot.OneNetAPI,
	}
	cloud := iot.NewCloud(oneNetCfg)
	if cloud == nil {
		log.Printf("[gpstracker]failed to new OneNet iot cloud")
		return
	}
	t := &gpsTracker{
		gps:    gps,
		logger: logger,
		cloud:  cloud,
	}

	util.WaitQuit(t.close)
	t.start()
}

type gpsTracker struct {
	gps    *dev.GPS
	cloud  iot.Cloud
	logger *util.GPSLogger
}

func (t *gpsTracker) start() {
	log.Printf("[gpstracker]start working")
	for {
		time.Sleep(2 * time.Second)
		pt, err := t.gps.Loc()
		if err != nil {
			log.Printf("[gpstracker]failed to get gps locations: %v", err)
			continue
		}
		t.logger.AddPoint(pt)
		v := &iot.Value{
			Device: "gps",
			Value:  pt,
		}
		go t.cloud.Push(v)
	}
}

func (t *gpsTracker) close() {
	t.gps.Close()
	t.logger.Close()
}
