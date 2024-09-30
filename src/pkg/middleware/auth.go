package middleware

import (
	"bytes"
	"hisaab-kitaab/pkg/config"
	"hisaab-kitaab/pkg/logger"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	kjwt "github.com/kataras/jwt"
	"go.uber.org/zap"
)

var (
	sep = []byte(".")
)

func joinParts(parts ...[]byte) []byte {
	return bytes.Join(parts, sep)
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			// xlength     string
			// userId      string
			accessToken string
		)
		alg := kjwt.EdDSA
		publicKey, err := kjwt.LoadPublicKeyEdDSA("./middlewares/id_rsa.pub")
		if err != nil {
			logger.Log(c).Error("io error opening public key file", zap.Error(err))
		}
		// Access Token will be appended with Bearer, need to get only the token
		accessTokenString := c.Request.Header.Get("Authorization")
		if accessTokenString == "" || !strings.Contains(accessTokenString, "Bearer") {
			logger.Log(c).Debug("accessTokenString:", zap.Any("Auth", accessTokenString))
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		accessToken = strings.Split(accessTokenString, " ")[1]
		/* not decrpyting the jwt token */
		// decryptedToken, err := decryptAccessToken(accessToken)
		// if err != nil {
		// 	logger.Log(c).Error("invalid value for encrypted access token", zap.Error(err))
		// 	//Not returning after receiving error, because it might be the case token was not encrypted in that case, we need to use that as it is. Return statement will be added in latter stages when it has been fully integrated with FE
		// } else {
		// 	accessToken = decryptedToken
		// }
		logger.Log(c).Debug("access token", zap.Any("Access Token", accessToken))
		c.Set(config.AccessToken, accessToken)

		accessTokenSplit := bytes.Split([]byte(accessToken), sep)
		headerEncoded := accessTokenSplit[0]
		payloadEncoded := accessTokenSplit[1]
		signatureEncoded := accessTokenSplit[2]
		signatureDecoded, _ := kjwt.Base64Decode(signatureEncoded)
		// Initialize a new instance of `Claims`
		claims := jwt.MapClaims{}

		tkn, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})

		if !tkn.Valid {
			logger.Log(c).Error("error", zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "error": "invalid auth token"})
			c.Abort()
			return
		}

		err = alg.Verify(publicKey, joinParts(headerEncoded, payloadEncoded), signatureDecoded)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				logger.Log(c).Error("error", zap.Error(err))
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "error": "invalid auth token"})
				c.Abort()
				return
			}
			logger.Log(c).Error("error", zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "error": "invalid auth token"})
			c.Abort()
			return
		}

		//check for claims before consuming
		if len(claims) == 0 {
			logger.Log().Error("claims are unavailable from the token")
			c.JSON(http.StatusForbidden, gin.H{"message": "Unauthorized", "error": "auth token invalid"})
			c.Abort()
			return
		}

		// checking for user_id, ucc, X-Length and scope is present in claims
		if claims["user_id"] == nil || claims["user_id"].(string) == "" || claims["user"] == nil || claims["user"].(string) == "" || claims["X-Length"] == nil || claims["X-Length"].(string) == "" {
			log.Println("user details is missing in token")
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "error": "invalid auth token"})
			c.Abort()
			return
		}

		authXLength := c.GetHeader("X-Length")
		log.Println("authXLength", authXLength, "claimsXLength", claims["X-Length"].(string))
		if authXLength != claims["X-Length"].(string) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "error": "invalid auth token"})
			c.Abort()
			return
		}

		c.Set(config.TOKEN, tkn)
		// userId = claims["user_id"].(string)
		// c.Set(config.USERID, userId)

		// xlength = claims["X-Length"].(string)
		// c.Set(config.XLENGTH, xlength)

		log.Println("claims", claims)

		requestID := uuid.New().String()
		c.Set(config.REQUESTID, requestID)

		logger.Log(c).Info("CALL STARTED", zap.Any("token", tkn))
		c.Next()
	}
}

// func decryptAccessToken(encryptedToken string) (string, error) {
// 	var (
// 		token string
// 	)

// 	encryptionKey := config.GetConfig().GetString("aes.secretkey256")
// 	cipher := utils.NewAesCipherService(encryptionKey, true)

// 	token, err := cipher.AuthTokenDecryption(encryptedToken)
// 	if err != nil {
// 		return token, err
// 	}

// 	return token, nil
// }

func CustomLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		if c.FullPath() != "/health" {
			latency := time.Since(start).Milliseconds()
			userID := c.GetString(config.USERID)
			uID := c.GetString(config.REQUESTID)
			// ucc := c.GetString(config.UCC)
			logger.Info(path,
				zap.String("requestID", uID),
				zap.String("leadId", "ucc"),
				zap.String("userId", userID),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Int64("latency", latency),
			)
		}
	}
}
