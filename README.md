# genomelink-go
[GENOME LINK](ttps://genomelink.io/developers) SDK, implemented in Go


# Install
```bash
$ go get -u github.com/ikasamah/genomelink-go
```

## Run example
```bash
$ go run $GOPATH/src/github.com/ikasamah/genomelink-go/example/main.go
Report: Genetic eye color 0 Tend to not have brown eyes
Sequence: "CACTAAGCACACAGAGAATAATGTCTAGAATCTGAGTGCCATGTTATCAAATTGTA..."
```

## OAuth Example
See [Official Document](https://genomelink.io/developers/docs/tutorial-oauth-example/) first.

`` bash
$ go run $GOPATH/src/github.com/ikasamah/genomelink-go/example/oauth.go -client_id <CLIENT_ID> -client_secret <CLIENT_SECRET>
```

Then, access http://localhost:8080/ and click auth link to get your personal token.

