GENOME LINK Go
==============
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[godocs]: https://godoc.org/github.com/ikasamah/genomelink-go/v1

[GENOME LINK](https://genomelink.io/developers) SDK, implemented in Go

# Install
```bash
$ go get -u github.com/ikasamah/genomelink-go
```

# Usage
```golang
import "github.com/ikasamah/genomelink-go/v1"

ctx := context.Background()
cli := genomelink.DefaultClient(ctx, "<YOUR_TOKEN>")

report, _ := cli.Report(ctx, genomelink.PhenotypeNameGeneticEyeColor, genomelink.PopulationEuropean)

fmt.Println(report.Summary.Text) // Tend to not have brown eyes
```

## Run example
```bash
$ go run $GOPATH/src/github.com/ikasamah/genomelink-go/example/main.go
Report: Genetic eye color 0 Tend to not have brown eyes
Sequence: "CACTAAGCACACAGAGAATAATGTCTAGAATCTGAGTGCCATGTTATCAAATTGTA..."
```

## OAuth Example
See [Official Document](https://genomelink.io/developers/docs/tutorial-oauth-example/) first, and register your app.

Run following.

```bash
$ go run $GOPATH/src/github.com/ikasamah/genomelink-go/example/oauth.go -client_id <CLIENT_ID> -client_secret <CLIENT_SECRET>
```

Then, access http://localhost:8080/ and click auth link to get your personal token.

