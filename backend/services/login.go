package services

import (
	"backend/clients"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func Login(username string, password string) (string, error) {
	user, err := clients.SelectUser(username)
	if err != nil {
		return "", err
	}

	passwordHash := hashMD5(password)
	if username != user.NombreUsuario || passwordHash != user.Contrasena {
		return "", errors.New("invalid credentials")
	}

	token, err := generateJWT(username)
	if err != nil {
		return "", errors.New("error generating token")
	}

	return token, nil
}

func hashMD5(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

var jwtSecret = []byte("your_secret_key") // Replace with your actual secret key

func generateJWT(username string) (string, error) {
	// Create the claims
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Token expiry time (72 hours)
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	return token.SignedString(jwtSecret)
}

/*var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Password string `json:"password"`
	Username    string `json:"email"`
}

type Claims struct {
	Username string `json:"email"`
	jwt.RegisteredClaims
}

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username    string `json:"email"`
	Password string `json:"-"`
}

type LoginService struct {
	DB *gorm.DB
}

func (s *LoginService) Login(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := s.DB.Where("email = ?", creds.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	expirationTime := jwt.NewNumericDate(time.Now().Add(5 * time.Minute))
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
*/

/*func Login(email string, password string) (string, error) {
	if strings.TrimSpace(email) == "" {
		return "", errors.New("email is required")
	}

	if strings.TrimSpace(password) == "" {
		return "", errors.New("password is required")
	}

	hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))

	userDAO, err := clients.SelectUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("error getting user from DB: %w", err)
	}

	if hash != userDAO.Password {
		return "", fmt.Errorf("invalid credentials")
	}

	// Generate JWT token
	token, err := generateJWT(email)
	if err != nil {
		return "", fmt.Errorf("error generating JWT token: %w", err)
	}

	return token, nil
}
*/
