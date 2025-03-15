package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/user"

	"github.com/zcalusic/sysinfo"
)

type SystemInfo struct {
	OS     OS     `json:"os"`
	Kernel Kernel `json:"kernel"`
	CPU    CPU    `json:"cpu"`
	Memory Memory `json:"memory"`
}

type OS struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Release string `json:"release"`
}

type Kernel struct {
	Release      string `json:"release"`
	Version      string `json:"version"`
	Architecture string `json:"architecture,omitempty"`
}

type CPU struct {
	Vendor  string `json:"vendor,omitempty"`
	Model   string `json:"model,omitempty"`
	Speed   uint   `json:"speed,omitempty"`   // CPU clock rate in MHz
	Cache   uint   `json:"cache,omitempty"`   // CPU cache size in KB
	Cpus    uint   `json:"cpus,omitempty"`    // number of physical CPUs
	Cores   uint   `json:"cores,omitempty"`   // number of physical CPU cores
	Threads uint   `json:"threads,omitempty"` // number of logical (HT) CPU cores
}

type Memory struct {
	Size uint `json:"size"` // Total physical memory in bytes
}

func main() {
	current, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if current.Uid != "0" {
		log.Fatal("requires superuser privilege")
	}

	fmt.Println("Getting system information...")
	var si sysinfo.SysInfo
	si.GetSysInfo()

	// Debug: Print raw sysinfo
	fmt.Println("\nRaw sysinfo:")
	rawData, _ := json.MarshalIndent(&si, "", "  ")
	fmt.Println(string(rawData))

	// Create our custom SystemInfo struct
	sysInfo := SystemInfo{
		OS: OS{
			Name:    si.OS.Name,
			Version: si.OS.Version,
			Release: si.OS.Release,
		},
		Kernel: Kernel{
			Release:      si.Kernel.Release,
			Version:      si.Kernel.Version,
			Architecture: si.Kernel.Architecture,
		},
		CPU: CPU{
			Vendor:  si.CPU.Vendor,
			Model:   si.CPU.Model,
			Speed:   si.CPU.Speed,
			Cache:   si.CPU.Cache,
			Cpus:    si.CPU.Cpus,
			Cores:   si.CPU.Cores,
			Threads: si.CPU.Threads,
		},
		Memory: Memory{
			Size: si.Memory.Size,
		},
	}

	fmt.Println("\nProcessed SystemInfo:")
	data, err := json.MarshalIndent(&sysInfo, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
