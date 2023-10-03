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
    var icon string
		if err != nil {
      fmt.Fprintf(os.Stderr, err.Error())
      goto error
		}
    capacity, err = strconv.Atoi(strings.TrimSpace(string(bs)))
    if err != nil {
      fmt.Fprintln(os.Stderr, err.Error())
      goto error
    }
    if capacity >= 90 {
      icon = "\uf240"
    } else if capacity >= 75 {
      icon = "\uf241"
    } else if capacity >= 50 {
      icon = "\uf242"
    } else if capacity >= 25 {
      icon = "\uf243"
    } else {
      stat.Urgent = true
      icon = "\uf244"
    }
    stat.FullText = fmt.Sprintf("%s %d%%", icon, capacity)
    goto next
  error:
			stat.Urgent = true
			stat.FullText = "battery"
	next:
		channel <- stat
		time.Sleep(10 * time.Second)
	}
}
