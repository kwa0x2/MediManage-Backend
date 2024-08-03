package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/MediManage-Backend/controllers"
)

func DistrictRoute(router *gin.Engine, districtController *controllers.DistrictController) {
	districtRoutes := router.Group("/api/v1/district")
	{
		districtRoutes.GET("", districtController.GetAll)
		districtRoutes.GET("/:province_name", districtController.GetAllByProvince)

	}
}

func ProvinceRoute(router *gin.Engine, provinceController *controllers.ProvinceController) {
	provinceRoutes := router.Group("/api/v1/province")
	{
		provinceRoutes.GET("", provinceController.GetAll)

	}
}

func AuthRoute(router *gin.Engine, authController *controllers.AuthController) {
	authRoutes := router.Group("/api/v1/auth")
	{
		authRoutes.POST("register", authController.Register)
		authRoutes.POST("login", authController.Login)
	}
}
