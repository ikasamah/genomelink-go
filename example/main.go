package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/ikasamah/genomelink-go/v1"
)

func main() {
	ctx := context.Background()
	cli := genomelink.DefaultClient(ctx, "GENOMELINKTEST001")
	//li.Logger, _ = zap.NewDevelopment()

	report, err := cli.Report(ctx, genomelink.PhenotypeNameGeneticEyeColor, genomelink.PopulationEuropean)

	if err != nil {
		fmt.Printf("Failed to get report: %s", err)
		return
	}

	fmt.Printf("Report: %s, %d, %s\n", report.Phenotype.DisplayName, report.Summary.Score, report.Summary.Text)

	sequence, err := cli.Sequence(ctx, genomelink.Region{
		Chromosome: genomelink.Chromosome1,
		Start:      100000,
		End:        100500,
	})

	if err != nil {
		fmt.Printf("Failed to get sequence: %s", err)
		return
	}

	s, _ := ioutil.ReadAll(sequence)
	fmt.Printf("Sequence: %s\n", s)
}
