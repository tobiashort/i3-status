package monitor

import (
	"net"
	"time"

	"github.com/t-hg/i3-status/def"
)

func InternalIP(channel chan def.Status) {
  for {
    stat := def.DefaultStatus()
    stat.Name = "internal_ip"
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
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


  
