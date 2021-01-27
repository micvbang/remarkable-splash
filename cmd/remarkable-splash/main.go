package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/disintegration/imaging"
	"github.com/micvbang/remarkable-splash/internal/splash"
)

func main() {
	flags := parseFlags()

	wifiConnected := make(chan struct{}, 1)
	go splash.WifiSignalConnected(wifiConnected)

	t0 := time.Now()

	for {
		select {
		case <-wifiConnected:
		case <-time.After(flags.pollInterval):
		}

		if time.Since(t0) < flags.pollInterval {
			continue
		}

		t0 = time.Now()

		img, err := splash.FetchNewest()
		if err != nil {
			fmt.Printf("failed to fetch image: %s\n", err)
			os.Exit(1)
		}

		img, err = splash.Resize(img)
		if err != nil {
			fmt.Printf("failed to resize image: %s\n", err)
		}

		err = imaging.Save(img, flags.output)
		if err != nil {
			log.Printf("Failed to save image to %s: %s", flags.output, err)
			os.Exit(1)
		}

	}
}

type flags struct {
	output       string
	pollInterval time.Duration
}

func parseFlags() flags {
	flags := flags{}

	flag.StringVar(&flags.output, "output", "", "path to output downloaded image")
	flag.DurationVar(&flags.pollInterval, "poll-interval", 6*time.Hour, "time between polling xkcd")

	flag.Parse()

	if flags.output == "" {
		flag.Usage()
		os.Exit(1)
	}

	return flags
}
