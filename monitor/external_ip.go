package monitor

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/t-hg/i3-status/def"
)

func ExternalIP(channel chan def.Status) {
  for {
    stat := def.DefaultStatus()
    stat.Name = "external_ip"
    client := http.Client{Timeout: 10*time.Second}
    res, err := client.Get("https://ifconfig.me")
    var bs []byte
    if err != nil {
      fmt.Fprintln(os.Stderr, err.Error())
      stat.Urgent = true
      stat.FullText = stat.Name
      goto next
    } 
    bs, err = io.ReadAll(res.Body)
    if err != nil {
      fmt.Fprintln(os.Stderr, err.Error())
      stat.Urgent = true
      stat.FullText = stat.Name
      goto next
    }
    stat.FullText = strings.TrimSpace(string(bs))
next:
    channel <- stat
    time.Sleep(10*time.Second)
  }
}


  
