package dev

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/jakefau/rpi-devices/util/geo"
)

// MockGPS ...
type MockGPS struct {
	index  int
	points []*geo.Point
}

// NewMockGPS ...
func NewMockGPS(csv string) *MockGPS {
	m := &MockGPS{}
	if err := m.open(csv); err != nil {
		return nil
	}
	return m
}

// Loc ...
func (m *MockGPS) Loc() (*geo.Point, error) {
	n := len(m.points)
	if n == 0 {
		return nil, errors.New("without data")
	}
	if m.index >= len(m.points) {
		m.index = 0
	}
	pt := m.points[m.index]
	m.index++
	return pt, nil
}

// Close ...
func (m *MockGPS) Close() {
	return
}

func (m *MockGPS) open(csv string) error {
	file, err := os.Open(csv)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		var timestamp string
		var lat, lon float64
		if _, err := fmt.Sscanf(line, "%19s,%f,%f\n", &timestamp, &lat, &lon); err != nil {
			return err
		}
		pt := &geo.Point{
			Lat: lat,
			Lon: lon,
		}
		m.points = append(m.points, pt)
	}
	file.Close()
	m.index = 0

	return nil
}
