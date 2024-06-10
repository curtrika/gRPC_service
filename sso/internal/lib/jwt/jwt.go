package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"gRPC_service/sso/internal/domain/models"
)

// NewToken creates new JWT token for given user and app.
func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
	//TODO: написать тесты
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Clams.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["app_id"] = app.ID

	// Подписываем токен, используя секретный ключ приложения
	tokenString, err := token.SignedString([]byte(app.Secret))
	if ee != nil {
		return "", err
	}

	return tokenString, nil
}
