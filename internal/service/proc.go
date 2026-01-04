package service

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/tousart/topz/internal/models"
)

type ProcService struct {
}

func NewProcService() *ProcService {
	return &ProcService{}
}

func (ps *ProcService) GetProc() ([]models.Proc, error) {
	pids, err := getProcPids()
	if err != nil {
		return nil, fmt.Errorf("/service: %v", err)
	}

	memTotal, err := getMemTotal()
	if err != nil {
		return nil, fmt.Errorf("/service: %v", err)
	}

	procs := make([]models.Proc, 0)

	procTimesBefore := make([]uint64, len(pids))
	for i, pid := range pids {
		procTime, err := getProcTime(pid)

		if err != nil {
			log.Printf("/service: %v\n", err)
		}

		procTimesBefore[i] = procTime
	}

	time.Sleep(models.GetProcCPUInterval * time.Second)

	for i, pid := range pids {
		procTimeBefore := procTimesBefore[i]

		procTimeAfter, err := getProcTime(pid)
		if err != nil {
			log.Printf("/service: %v\n", err)
			continue
		}

		cpuPercent := math.Round(float64(procTimeAfter-procTimeBefore)/models.TicksPerSecond*100.0*10.0) / 10.0

		memPercent := 0.0
		if memTotal > 0 {

			procMem, err := getProcMemory(pid)
			if err != nil {
				log.Printf("/service: %v\n", err)
				continue
			}

			memPercent = math.Round(float64(procMem)/float64(memTotal)*100.0*10.0) / 10.0
		}

		procs = append(procs,
			models.Proc{
				PID:        pid,
				CPUPercent: cpuPercent,
				MemPercent: memPercent,
			})
	}

	return procs, nil
}
