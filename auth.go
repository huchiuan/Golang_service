package main
import (
    "errors"
    "fmt"
    "github.com/golang-jwt/jwt"
    "time"
)
var jwtKey = []byte("3Aju5DBXtMaFjNiANXclJ3WTcdY2pdzpi6y1USGUH0cHCLTVjZyIEQVtZbupnDGaplHgJ4/akowsthFcxWfaTRBaZZOfrpS765t+227s1W+lFPb9Yt5oeVx+yn2yOAMMIx7pCXyDM58SzdGQPEOx8+uK0gyWYhkdVgRnJIzb4iOzR6h+7mnyPfqguo3Bz9kF1u4usZRcpf4mfgB0/+G/hDe9/Y6yiqM3HqDlKQ==")
type authClaims struct {
    jwt.StandardClaims
    UserID uint `json:"userId"`
}
func generateToken(user User) (string, error) {
    expiresAt := time.Now().Add(24 * time.Hour).Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
        StandardClaims: jwt.StandardClaims{
            Subject:   user.Username,
            ExpiresAt: expiresAt,
        },
        UserID: user.ID,
    })
    tokenString, err := token.SignedString(jwtKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

func validateToken(tokenString string) (uint, string, error) {
    var claims authClaims
    token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return jwtKey, nil
    })
    if err != nil {
        return 0, "", err
    }
    if !token.Valid {
        return 0, "", errors.New("invalid token")
    }
    id := claims.UserID
    username := claims.Subject
    return id, username, nil
}