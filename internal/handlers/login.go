package handlers

import (
	"net/http"

	"github.com/Yummiii/Nayu/internal/auth"
	"github.com/Yummiii/Nayu/internal/database"
	"github.com/Yummiii/Nayu/internal/discord"
	"github.com/gin-gonic/gin"
)


func LoginHandler() gin.HandlerFunc {
	oauthCfg := discord.GetOauthCfg()

	return func(c *gin.Context) {
		code := c.Query("code")

		if code == "" {
			url := oauthCfg.AuthCodeURL("state")
			c.Redirect(http.StatusTemporaryRedirect, url)
			return
		}

		token, err := oauthCfg.Exchange(c, code)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to exchange token: %v", err)
			return
		}

		discordUser, err := discord.GetUserInfo(token.AccessToken)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to get user info: %v", err)
			return
		}

		// db := database.GetDB()
		// dbUser, err := database.GetUserByDiscordID(db, discordUser.ID)

		// jwtToken, err := auth.GenerateJWT(user)
		// if err != nil {
		// 	c.String(http.StatusInternalServerError, "Failed to generate JWT: %v", err)
		// 	return
		// }

		// c.SetCookie("auth_token", jwtToken, 3600, "/", "", false, true)

		c.Redirect(http.StatusPermanentRedirect, "/")
	}
}
