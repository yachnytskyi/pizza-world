package handler

import (
	"pizza-world/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signIn)
		auth.POST("/sign-in", h.signUp)
	}

	api := router.Group("/api")
	{
		recipes := api.Group("/recipes")
		{
			recipes.GET("/", h.getAllRecipes)
			recipes.GET("/:id", h.getRecipeById)
			recipes.POST("/", h.createRecipe)
			recipes.PUT("/:id", h.updateRecipe)
			recipes.DELETE("/:id", h.deleteRecipe)

			ingredients := recipes.Group(":id/ingredients")
			{
				ingredients.GET("/", h.getAllIngredients)
				ingredients.GET("/:ingredient_id", h.getIngredientById)
				ingredients.POST("/", h.createIngredient)
				ingredients.PUT("/:ingredient_id", h.updateIngredient)
				ingredients.DELETE("/:ingredient_id", h.deleteIngredient)
			}
		}
	}

	return router
}
