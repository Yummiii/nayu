package discord

import "golang.org/x/oauth2"

func GetOauthCfg() oauth2.Config {
	return oauth2.Config{
		ClientID: "429746766917926914",
		ClientSecret: "H-qICKkvlQjLKldtqoedMTna_-qYtpvz",
		Endpoint: oauth2.Endpoint{
			AuthURL: "https://discord.com/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
		RedirectURL: "http://localhost:8080/login",
		Scopes:      []string{"identify"},
	}
}
