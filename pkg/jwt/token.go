package jwt

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"time"
)

type NewClaims struct {
	UserID uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	jwt.RegisteredClaims
}

var (
	//AccessTokenExpired  = os.Getenv("ACCESS_TOKEN_EXPIRED")
	//RefreshTokenExpired = os.Getenv("REFRESH_TOKEN_EXPIRED")

	//Secret = os.Getenv("JWT_SECRET")
	Secret = "secret"
)

func GenerateToken(user models.User, jwtID string, expired int) (string, error) {

	// prepare the object Claims
	claims := NewClaims{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jwtID,
			Issuer:    "patungan.cuy",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expired))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	logrus.Println("token : ", token)

	// make the accessToken more secure with the jwt_secret
	tokenSigned, err := token.SignedString([]byte(Secret))
	logrus.Println("tokenSigned : ", tokenSigned)
	if err != nil {
		logrus.Printf("cannot signed jwt access token string")
		return "", err
	}

	return tokenSigned, nil
}

func ValidateToken(tokenString string) (*NewClaims, error) {

	// parse tokenString (with jwt secret) to get the real token (without jwt secret)
	token, err := jwt.ParseWithClaims(tokenString, &NewClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(Secret), nil
	})

	if err != nil {
		logrus.Printf("error validate token")
		return nil, err
	}

	// parse token to Claims struct, if ok, return it
	if claims, ok := token.Claims.(*NewClaims); ok && token.Valid {
		logrus.Println("claims : ", claims)
		return claims, nil
	} else {
		logrus.Error("Failed to convert claims to NewClaims")
		return nil, fmt.Errorf("invalid token claims")
	}

}
