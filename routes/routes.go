package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/MediManage-Backend/controllers"
	"github.com/kwa0x2/MediManage-Backend/middlewares"
)

func DistrictRoute(router *gin.Engine, districtController *controllers.DistrictController) {
	districtRoutes := router.Group("/api/v1/district")
	{
		districtRoutes.GET("", districtController.GetAll)
		districtRoutes.GET("/:provinceName", districtController.GetAllByProvince)

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
		authRoutes.POST("logout", authController.Logout)
		authRoutes.GET("", middlewares.SessionMiddleware(), authController.CheckAuth)

	}
}

func UserRoute(router *gin.Engine, userController *controllers.UserController) {
	userRoutes := router.Group("/api/v1/user")
	userRoutes.Use(middlewares.SessionMiddleware())
	{
		userRoutes.GET("", userController.GetAll)
		userRoutes.POST("", userController.Create)
		userRoutes.PUT("", userController.Update)
		userRoutes.DELETE("/:userId", userController.Delete)
	}
}

func EmployeeRoute(router *gin.Engine, employeeController *controllers.EmployeeController) {
	employeeRoutes := router.Group("/api/v1/employee")
	employeeRoutes.Use(middlewares.SessionMiddleware())
	{
		employeeRoutes.GET("", employeeController.GetAll)
		employeeRoutes.POST("withworkday", employeeController.CreateWithWorkDay)
		employeeRoutes.PUT("", employeeController.Update)
		employeeRoutes.DELETE("/:employeeId", employeeController.Delete)
	}
}

func ClinicRoute(router *gin.Engine, clinicController *controllers.ClinicController) {
	clinicRoute := router.Group("/api/v1/clinic")
	clinicRoute.Use(middlewares.SessionMiddleware())
	{
		clinicRoute.GET("", clinicController.GetAll)
		clinicRoute.GET("/hospital", clinicController.GetAllHospitalClinic)
		clinicRoute.POST("/hospital", clinicController.CreateHospitalClinic)
		clinicRoute.DELETE("/hospital/:clinicName", clinicController.DeleteHospitalClinic)
		clinicRoute.PUT("/hospital", clinicController.Update)

	}
}
func JobGroupRoute(router *gin.Engine, jobGroupController *controllers.JobGroupController) {
	jobGroupRoutes := router.Group("/api/v1/jobgroup")
	jobGroupRoutes.Use(middlewares.SessionMiddleware())
	{
		jobGroupRoutes.GET("", jobGroupController.GetAll)
	}
}
func TitleRoute(router *gin.Engine, titleController *controllers.TitleController) {
	titleRoutes := router.Group("/api/v1/title")
	titleRoutes.Use(middlewares.SessionMiddleware())
	{
		titleRoutes.GET("/:jobGroupName", titleController.GetAllTitleByJobGroupName)
	}
}

func HospitalRoute(router *gin.Engine, hospitalController *controllers.HospitalController) {
	hospitalRoutes := router.Group("/api/v1/hospital")
	hospitalRoutes.Use(middlewares.SessionMiddleware())
	{
	}
}
