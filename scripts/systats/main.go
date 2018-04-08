package main

import (
    "fmt"
    "time"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/cpu"
)

func main() {
    m, _ := mem.VirtualMemory()

    i, _ := time.ParseDuration("2s")
    c, _ := cpu.Percent(i, false)

    fmt.Printf("cpu %.0f%% mem %.0f%%", c[0], m.UsedPercent)
}
