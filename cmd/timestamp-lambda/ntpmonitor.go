package main

import (
	"fmt"
	"github.com/sigstore/timestamp-authority/pkg/ntpmonitor"
)

var ntpm *ntpmonitor.NTPMonitor

func startNtpMonitor() {

	// Load external file here.
	ntpMonitoring := "" // viper.GetString("ntp-monitoring")

	fmt.Println("NTP Monitor: Starting")
	ntpm, err := ntpmonitor.New(ntpMonitoring)
	if err != nil { panic(err) }

	ntpm.Start()
	fmt.Println("NTP Monitor: Started")	

}

func stopNtpMonitor() {
	if ntpm != nil {
		fmt.Println("NTP Monitor: Stopping")
		ntpm.Stop()
		fmt.Println("NTP Monitor: Stopped")
	}
}
