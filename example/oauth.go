package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/oauth2"
)

var endpopint = oauth2.Endpoint{
	AuthURL:  "https://genomelink.io/oauth/authorize",
	TokenURL: "https://genomelink.io/oauth/token",
}

var state = strconv.FormatUint(rand.Uint64(), 36)

func main() {
	port := flag.Int("port", 8080, "Listen port")
	clientID := flag.String("client_id", os.Getenv("GENOMELINK_CLIENT_ID"), "client id, you may override by GENOMELINK_CLIENT_ID env ver.")
	clientSecret := flag.String("client_secret", os.Getenv("GENOMELINK_CLIENT_SECRET"), "client secret, you may override by GENOMELINK_CLIENT_SECRET env ver.")
	flag.Parse()

	config := oauth2.Config{
		ClientID:     *clientID,
		ClientSecret: *clientSecret,
		Endpoint:     endpopint,
		RedirectURL:  fmt.Sprintf("http://localhost:%d/callback", *port),
		Scopes:       []string{"report:eye-color", "report:beard-thickness", "report:morning-person"},
	}

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/auth", authHandler(config))
	http.HandleFunc("/callback", callbackHandler(config))

	addr := fmt.Sprintf(":%d", *port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	html := `<html><body><a href="/auth">Log in with GENOMELINK</a></body></html>`
	fmt.Fprint(w, html)
}

func authHandler(config oauth2.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := config.AuthCodeURL(state)
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}

func callbackHandler(config oauth2.Config) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("state") != state {
			log.Printf("[ERROR] Invalid state")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token, err := config.Exchange(context.Background(), r.FormValue("code"))
		if err != nil {
			log.Printf("[ERROR] Failed to exchange: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tpl := `<html><body><ul>
			<li>token_type: %s</li>
			<li>access_token: %s</li>
			<li>refresh_token: %s</li>
			<li>expiry: %s</li>
			<li>scope: %s</li>
		</ul></body></html>`

		html := fmt.Sprintf(tpl, token.TokenType, token.AccessToken, token.RefreshToken, token.Expiry, token.Extra("scope"))
		fmt.Fprint(w, html)
	}
}
