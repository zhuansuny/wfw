# wfw
go web framework
#### wfw 是一个简单的web 框架

快速开始
```
	r := wfw.New()
	r.Use(wfw.Logger())
	api := r.Group("/api")
	api.GET("/test",func(c *wfw.Context) {
		c.JSON(http.StatusOK, "Hello world")
	})

	r.Run(":9999")
```
