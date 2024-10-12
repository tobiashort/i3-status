package main

import (
	"encoding/json"
	"fmt"

	"github.com/tobiashort/i3-status/def"
	"github.com/tobiashort/i3-status/monitor"
	"github.com/tobiashort/i3-status/must"
)

func toJsonString(v interface{}) string {
	return string(must.Do2(json.Marshal(v)))
}

func main() {
	header := def.Header{Version: 1}
	fmt.Println(toJsonString(header))
	fmt.Println("[")
	fmt.Println("  []")

	stats := make([]def.Status, 4)
	datetimeStatChan := make(chan def.Status)
	externalIPStatChan := make(chan def.Status)
	internalIPStatChan := make(chan def.Status)
	batteryStatChan := make(chan def.Status)

	go monitor.Datetime(datetimeStatChan)
	go monitor.ExternalIP(externalIPStatChan)
	go monitor.InternalIP(internalIPStatChan)
	go monitor.Battery(batteryStatChan)

	for {
		select {
		case stat := <-datetimeStatChan:
			stats[3] = stat.Invert()
		case stat := <-internalIPStatChan:
			stats[2] = stat
		case stat := <-externalIPStatChan:
			stats[1] = stat.Invert()
		case stat := <-batteryStatChan:
			stats[0] = stat
		}
		fmt.Println(",", toJsonString(stats))
	}
}
