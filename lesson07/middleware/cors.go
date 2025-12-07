//这cors全靠AI神力，真不会（orz）

package middleware

import "github.com/gin-gonic/gin"

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 【最关键】告诉浏览器：“允许来自任何来源的请求”（开发模式用）
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 如果前端需要发送Cookie或Authorization头，这里不能用*，必须指定明确来源，例如：
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

		// 2. 告诉浏览器：“我允许请求携带这些额外的头信息”
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		// `Authorization`头（放JWT Token的）和`Content-Type`头（说明数据是JSON）通常必须有。

		// 3. 告诉浏览器：“我允许这些HTTP方法”
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 这涵盖了RESTful API的基本方法。

		// 4. 【处理“预检请求”】—— 这是CORS的独特机制
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // 直接返回“成功，但无内容”
			return
		}
		// 当浏览器发送非简单请求（如带JSON的POST）前，会自动先发一个OPTIONS请求来“探路”。
		// 你的服务器只需返回这些CORS头，并响应204即可，不需要执行业务逻辑。

		// 5. 如果不是OPTIONS请求，就继续走后面的路由和业务逻辑（登录、查询等）
		c.Next()
	}
}

//🎯 CORS 是干什么的？
//想象一下：你的前端项目（比如用Vue/React写的页面）运行在 http://localhost:3000，而你的Go后端运行在 http://localhost:8080。
//
//对于浏览器来说，端口不同（3000 vs 8080）就是不同的“源”（Origin）。
//
//浏览器出于安全考虑，默认禁止一个“源”的网页向另一个“源”的服务器发起请求。这就是 “同源策略”。
//
//CORS 机制，就是由后端明确地告诉浏览器：“我允许来自 http://localhost:3000 的请求过来，你放行吧！”。
//
//你提供的这段代码，就是你的Go服务器在响应里对浏览器喊的那几句话。
