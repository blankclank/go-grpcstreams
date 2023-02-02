package main

import (
	"log"
	"time"

	hardwaremonitoring "grpcstreams/proto"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

type Server struct {
	hardwaremonitoring.UnimplementedHardwareMonitorServer
}

func (s *Server) Monitor(req *hardwaremonitoring.EmptyRequest, stream hardwaremonitoring.HardwareMonitor_MonitorServer) error {
	timer := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-timer.C:
			hwStats, err := s.GetStats()
			if err != nil {
				log.Println(err.Error())
			} else {

			}
			err = stream.Send(hwStats)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}

func (s *Server) GetStats() (*hardwaremonitoring.HardwareStats, error) {

	mem, err := memory.Get()
	if err != nil {
		return nil, err
	}
	// Extract CPU stats
	cpu, err := cpu.Get()
	if err != nil {
		return nil, err
	}
	// Create our response object
	hwStats := &hardwaremonitoring.HardwareStats{
		Cpu:        int32(cpu.Total),
		MemoryFree: int32(mem.Free),
		MemoryUsed: int32(mem.Used),
	}

	return hwStats, nil
}
