package service

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getProcPids() ([]int, error) {
	files, err := os.ReadDir("/proc")
	if err != nil {
		return nil, fmt.Errorf("GetPids: %v", err)
	}

	pids := make([]int, 0)

	for _, file := range files {
		if file.IsDir() {
			pid, err := strconv.Atoi(file.Name())
			if err != nil {
				continue
			}

			pids = append(pids, pid)
		}
	}

	if len(pids) == 0 {
		return nil, errors.New("pid directories not found")
	}

	return pids, nil
}

func getProcTime(pid int) (uint64, error) {
	data, err := os.ReadFile(fmt.Sprintf("/proc/%d/stat", pid))
	if err != nil {
		return 0, fmt.Errorf("readProcTime: %v", err)
	}

	stats := strings.Split(string(data), " ")

	// proc time in user mode
	utime, err := strconv.ParseUint(stats[13], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("readProcTime: %v", err)
	}

	// proc time in kernel mode
	stime, err := strconv.ParseUint(stats[14], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("readProcTime: %v", err)
	}

	return utime + stime, nil
}

func getProcMemory(pid int) (uint64, error) {
	data, err := os.ReadFile(fmt.Sprintf("/proc/%d/status", pid))
	if err != nil {
		return 0, fmt.Errorf("getProcMemory: %v", err)
	}

	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {
		if after, ok := strings.CutPrefix(line, "VmRSS:"); ok {

			memValue, err := strconv.ParseUint(
				strings.TrimSpace(strings.TrimSuffix(after, "kB")), 10, 64)

			if err != nil {
				return 0, fmt.Errorf("getProcMemory: %v", err)
			}

			// kB
			return memValue, nil
		}
	}

	return 0, errors.New("getProcMemory: VmRSS not found")
}

func getMemTotal() (uint64, error) {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return 0, fmt.Errorf("getMemTotal: %v", err)
	}

	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {
		if after, ok := strings.CutPrefix(line, "MemTotal:"); ok {

			memTotalValue, err := strconv.ParseUint(
				strings.TrimSpace(strings.TrimSuffix(after, "kB")), 10, 64)

			if err != nil {
				return 0, fmt.Errorf("getMemTotal: %v", err)
			}

			// kB
			return memTotalValue, nil
		}
	}

	return 0, errors.New("getMemTotal: MemTotal not found")
}
