package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/t-hg/i3-status/def"
	"github.com/t-hg/i3-status/monitor"
	"github.com/t-hg/i3-status/must"
)

func toJsonString(v interface{}) string {
  return string(must.Do2(json.Marshal(v)))
}

func main() {
  header := def.Header{Version: 1}
  fmt.Println(toJsonString(header))
  fmt.Println("[")
  fmt.Println("  []")

  stats := make([]def.Status, 3)
  datetimeStatChan := make(chan def.Status)
  externalIPStatChan := make(chan def.Status)
  internalIPStatChan := make(chan def.Status)

  go monitor.Datetime(datetimeStatChan)
  go monitor.ExternalIP(externalIPStatChan)
  go monitor.InternalIP(internalIPStatChan)

  go func() {
      for {
        select {
        case stat := <-datetimeStatChan:
          stats[2] = stat
        case stat := <-internalIPStatChan:
          stats[1] = stat
        case stat := <-externalIPStatChan:
          stats[0] = stat
        }
    }
  }()

  for {
    stats[1].Invert()
    fmt.Println(",", toJsonString(stats))
    time.Sleep(500*time.Millisecond)
  }
}
