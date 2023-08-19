package url

import (
	"strings"

	"github.com/golang-jwt/jwt"
)

var SECRET = []byte("take_me_to_this_link_please")

func HashURL(url string) (string, error) {
	if !strings.Contains(url, "http") || !strings.Contains(url, "https") {
		url = "http://" + url
	}

	slug := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"url": url,
	})

	return slug.SignedString([]byte(SECRET))
}

func VerifyURL(token string) (string, error) {
	slug, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})
	if err != nil {
		return "", err
	}

	url := slug.Claims.(jwt.MapClaims)["url"]

	return url.(string), nil
}
