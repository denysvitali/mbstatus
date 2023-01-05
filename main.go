package main

import (
	"encoding/json"
	"fmt"
	mb "github.com/denysvitali/go-mobilebroadband"
	"os"
	"strings"
)

func PrintErr(format string, args ...any) {
	_, _ = os.Stderr.WriteString(fmt.Sprintf(format, args...))
}

type Response struct {
	Imei               string               `json:"imei"`
	OperatorName       string               `json:"operatorName"`
	OperatorCode       string               `json:"operatorCode"`
	SignalByTechnology map[string]mb.Signal `json:"signal"`
}

func main() {
	b, err := mb.New()
	if err != nil {
		PrintErr("unable to create mb: %v", err)
		os.Exit(1)
	}

	modems, err := b.Modems()
	if err != nil {
		PrintErr("unable to get modems: %v", err)
		os.Exit(1)
	}

	if len(modems) == 0 {
		PrintErr("cannot find any modem")
		os.Exit(1)
	}

	modem := modems[0]
	err = modem.SetupPeriodicPolling(10)
	if err != nil {
		PrintErr("unable to setup periodic polling: %v", err)
		os.Exit(1)
	}

	signalByTechnology := map[string]mb.Signal{}

	for _, v := range []mb.Technology{
		mb.TechnologyEvdo, mb.TechnologyCdma,
		mb.TechnologyGsm, mb.TechnologyUmts,
		mb.TechnologyLte, mb.TechnologyNr5g,
	} {
		s, err := modem.GetSignal(v)
		if err != nil {
			PrintErr("unable to get signal for %s: %v", v, err)
			continue
		}
		if s == nil {
			continue
		}
		signalByTechnology[strings.ToLower(string(v))] = *s
	}

	statuses, err := b.Status()
	if err != nil {
		PrintErr("unable to get statuses: %v", err)
		os.Exit(1)
	}

	if len(statuses) == 0 {
		PrintErr("len(statuses) == 0")
		os.Exit(1)
	}

	// Let's take the first status
	status := statuses[0]

	r := Response{
		Imei:               status.M3gpp.Imei,
		OperatorName:       status.M3gpp.OperatorName,
		OperatorCode:       status.M3gpp.OperatorCode,
		SignalByTechnology: signalByTechnology,
	}

	// Setup Signal Polling

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(r)
	if err != nil {
		PrintErr("unable to encode JSON: %v", err)
		os.Exit(1)
	}
}
