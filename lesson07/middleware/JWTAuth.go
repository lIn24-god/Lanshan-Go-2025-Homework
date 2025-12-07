package middleware

import (
	"fmt"
	"lesson06/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth 认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从请求头获取Token，格式：Authorization: Bearer <token>
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "请求未携带token，无权访问"})
			return
		}

		// 2. 提取Token部分（去掉"Bearer "前缀）
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token格式错误"})
			return
		}
		tokenString := parts[1]

		// 3. 调用utils层解析和验证Token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			// Token无效或过期
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "无效的token或已过期"})
			return
		}

		// 4. Token验证通过，将当前用户名存入上下文，供后续处理函数使用
		c.Set("username", claims.Username)
		fmt.Printf("[JWTAuth] 用户 %s 通过认证\n", claims.Username)

		// 5. 放行，执行下一个中间件或最终的业务处理函数
		c.Next()
	}
}
