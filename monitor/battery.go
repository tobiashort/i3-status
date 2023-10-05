package monitor

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/t-hg/i3-status/def"
)

func Battery(channel chan def.Status) {
	for {
		stat := def.DefaultStatus()
		stat.Name = "battery"
		filename := "/sys/class/power_supply/BAT0/capacity"
		bs, err := os.ReadFile(filename)
    var capacity int
		if err != nil {
      fmt.Fprintf(os.Stderr, err.Error())
      goto error
		}
    capacity, err = strconv.Atoi(strings.TrimSpace(string(bs)))
    if err != nil {
      fmt.Fprintln(os.Stderr, err.Error())
      goto error
    }
    stat.FullText = fmt.Sprintf("\uf0e7%d%%", capacity)
    goto next
  error:
			stat.Urgent = true
			stat.FullText = "battery"
	next:
		channel <- stat
		time.Sleep(10 * time.Second)
	}
}
