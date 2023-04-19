package main

import (
	"fmt"
	"github.com/breedish/tbot/internal/bot/context"
	"os"

	"github.com/krzysztofreczek/go-structurizr/pkg/scraper"
	"github.com/krzysztofreczek/go-structurizr/pkg/view"
)

const (
	scraperConfig = "scraper.yml"
	viewConfig    = "view.yml"
	outputFile    = "out/view-%s.plantuml"
)

func main() {
	var gatewayApp, _, _ = context.NewApplication()
	scrape(gatewayApp, "bot")
}

func scrape(app interface{}, name string) {
	s, err := scraper.NewScraperFromConfigFile(scraperConfig)
	if err != nil {
		panic(err)
	}

	structure := s.Scrape(app)

	outFileName := fmt.Sprintf(outputFile, name)
	outFile, err := os.Create(outFileName)
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = outFile.Close()
	}()

	v, err := view.NewViewFromConfigFile(viewConfig)
	if err != nil {
		panic(err)
	}

	err = v.RenderStructureTo(structure, outFile)
	if err != nil {
		panic(err)
	}
}
