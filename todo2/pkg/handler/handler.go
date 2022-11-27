package handler

import (
	"github.com/Serminaz/GoRun/todo2/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct { // работа с  http
	// принимаем запросы от пользователя и передаем дальше в сервис
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}

}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New() // инициализация роутера

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp) // регистрация и авторизация
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity) // работа со списками задач
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList) // иницилизация и добавка методов
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
