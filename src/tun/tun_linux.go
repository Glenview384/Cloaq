//go:build linux

package tun

import (
	"os"

	"cloaq/src/tun/lintun"
)

type linuxDevice struct {
	name string
	f    *os.File
}

func (t *linuxDevice) Name() string                { return t.name }
func (t *linuxDevice) Start() error                { return nil }
func (t *linuxDevice) Close() error                { return t.f.Close() }
func (t *linuxDevice) Read(p []byte) (int, error)  { return t.f.Read(p) }
func (t *linuxDevice) Write(p []byte) (int, error) { return t.f.Write(p) }
func (t *linuxDevice) File() *os.File              { return t.f }

// InitDevice creates a L3 TUN on Linux
func InitDevice() (Device, error) {
	name := "cloaq0"
	f, err := lintun.CreateTUN(name)
	if err != nil {
		return nil, err
	}
	return &linuxDevice{name: name, f: f}, nil
}
