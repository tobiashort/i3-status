package monitor

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/tobiashort/i3-status/def"
)

func InternalIP(channel chan def.Status) {
  for {
    stat := def.DefaultStatus()
    stat.Name = "internal_ip"
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
      fmt.Fprintln(os.Stderr, err.Error())
      stat.Urgent = true
      stat.FullText = stat.Name
      goto next
    } 
    stat.FullText = conn.LocalAddr().(*net.UDPAddr).IP.String()
next:
    channel <- stat
    time.Sleep(10*time.Second)
  }
}


  
