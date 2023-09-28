package monitor

import (
	"time"

	"github.com/t-hg/i3-status/def"
)

func Datetime(channel chan def.Status) {
  for {
    stat := def.DefaultStatus()
    stat.Name = "datetime"
    stat.FullText = time.Now().Format("2006-01-02 15:04")
    channel <- stat
    time.Sleep(500*time.Millisecond)
  }
}
