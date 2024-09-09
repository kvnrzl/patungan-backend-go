package jwt

import (
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/models"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"time"
)

type Claims struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	jwt.RegisteredClaims
}

var (
	//AccessTokenExpired  = os.Getenv("ACCESS_TOKEN_EXPIRED")
	//RefreshTokenExpired = os.Getenv("REFRESH_TOKEN_EXPIRED")

	//Secret = os.Getenv("JWT_SECRET")
	Secret = "secret"
)

func GenerateToken(user models.User, expired int) (string, error) {

	// prepare the object Claims
	claims := Claims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "patungan.cuy",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expired))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	logrus.Println("token : ", token)

	// make the accessToken more secure with the jwt_secret
	fmt.Println("Secret : ", Secret)
	tokenSigned, err := token.SignedString([]byte(Secret))
	logrus.Println("tokenSigned : ", tokenSigned)
	if err != nil {
		logrus.Printf("cannot signed jwt access token string")
		return "", err
	}

	return tokenSigned, nil
}

func ValidateToken(tokenString string) (*Claims, error) {

	// parse tokenString (with jwt secret) to get the real token (without jwt secret)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// check dulu jika token methodnya sesuai dengan yang diinginkan, jika sesuai maka kembalikan secret key nya
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return Secret, nil
	})

	if err != nil {
		logrus.Printf("error validate token")
		return nil, err
	}

	// parse token to Claims struct, if ok, return it
	if claims, ok := token.Claims.(Claims); ok && token.Valid {
		return &claims, nil
	}

	return nil, err

}
