package discord

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

func GetUserInfo(accessToken string) (*User, error) {
	req, err := http.NewRequest("GET", "https://discord.com/api/users/@me", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (user User) AvatarURL() string {
	if user.Avatar == "" {
		return fmt.Sprintf("https://cdn.discordapp.com/embed/avatars/%d.png", (int(user.ID[0]) % 5))
	}
	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", user.ID, user.Avatar)
}
