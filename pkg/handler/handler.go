package handler


type Handler struct {

}


func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()


	auth := router.Group(relativePath: "/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group(relativePath:"/api")
	{
		lists := api.Group(relativePath:"/lists")
		{
			lists.POST(relativePath:"/", h.createList)
			lists.GET(relativePath:"/", h.getAllLists)
			lists.GET(relativePath:"/:id", h.getListById)
			lists.PUT(relativePath:"/:id", h.updateList)
			lists.DELETE(relativePath:"/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST(relativePath:"/", h.createItem)
				items.GET(relativePath:"/", h.getAllItems)
				items.GET("/:id", h.getItemById)
				items.PUT("/:id", h.updateItem)
				items.DELETE("/:id", h.deleteItem)
			}
		}

	}

	return router
}