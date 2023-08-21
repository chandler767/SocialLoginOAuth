package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/chandler767/SocialLoginOAuth/packages/dontlist"
	"github.com/chandler767/SocialLoginOAuth/packages/token"

	"github.com/gorilla/mux"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
)

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)
	var state = token.New("state")
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)
	return state
}

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "https://socialloginoauth-production.up.railway.app/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func main() {
	p := mux.NewRouter()
	p.HandleFunc("/", handleMain)
	p.HandleFunc("/login", handleGoogleLogin)
	p.HandleFunc("/callback", handleGoogleCallback)

	p.PathPrefix("/").Handler(http.FileServer(dontlist.DontListFiles{Fs: http.Dir("./static/")})) // Index file server.

	log.Println("listening on localhost:" + os.Getenv("PORT"))
	srv := &http.Server{
		Handler:      p,
		Addr:         "0.0.0.0:" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
