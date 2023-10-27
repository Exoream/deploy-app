package middlewares

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	godotenv.Load(".env")
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(os.Getenv("JWT_SECRET")),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId uuid.UUID, role string) (string, error) {
    claims := jwt.MapClaims{}
    claims["authorized"] = true
    claims["userId"] = userId.String()
    claims["role"] = role
    claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}


func ExtractToken(e echo.Context) uuid.UUID {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string) // Mengambil userID sebagai string UUID
		parsedID, err := uuid.Parse(userId)
		if err != nil {
			// Handle error
			return uuid.Nil
		}
		return parsedID
	}
	return uuid.Nil
}

func ExtractRole(c echo.Context) (string, error) {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    role, ok := claims["role"].(string)
    if !ok {
        return "", errors.New("role not found in token claims")
    }
    return role, nil
}

