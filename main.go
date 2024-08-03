package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kwa0x2/MediManage-Backend/cache"
	"github.com/kwa0x2/MediManage-Backend/config"
	"github.com/kwa0x2/MediManage-Backend/controllers"
	"github.com/kwa0x2/MediManage-Backend/repositories"
	"github.com/kwa0x2/MediManage-Backend/routes"
	"github.com/kwa0x2/MediManage-Backend/services"
	"log"
)

func main() {
	config.LoadEnv()
	router := gin.New()
	config.PostgreConnection()
	store := config.RedisSession()
	router.Use(sessions.Sessions("authorization", store))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	redisClient := config.NewRedisClient()

	districtCache := &cache.DistrictCache{RedisClient: redisClient}
	provinceCache := &cache.ProvinceCache{RedisClient: redisClient}

	districtRepository := &repositories.DistrictRepository{DB: config.DB}
	districtService := &services.DistrictService{DistrictRepository: districtRepository}
	districtController := &controllers.DistrictController{DistrictService: districtService, DistrictCache: districtCache}

	provinceRepository := &repositories.ProvinceRepository{DB: config.DB}
	provinceService := &services.ProvinceService{ProvinceRepository: provinceRepository}
	provinceController := &controllers.ProvinceController{ProvinceService: provinceService, ProvinceCache: provinceCache}

	userRepository := &repositories.UserRepository{DB: config.DB}
	userService := &services.UserService{UserRepository: userRepository}
	//userController := &controllers.ProvinceController{ProvinceService: provinceService}

	hospitalRepository := &repositories.HospitalRepository{DB: config.DB}
	//hospitalService := &services.HospitalService{HospitalRepository: hospitalRepository}
	//provinceController := &controllers.ProvinceController{ProvinceService: provinceService}

	authService := &services.AuthService{UserRepository: userRepository, HospitalRepository: hospitalRepository}
	authController := &controllers.AuthController{AuthService: authService, UserService: userService}

	routes.DistrictRoute(router, districtController)
	routes.ProvinceRoute(router, provinceController)
	routes.AuthRoute(router, authController)

	if err := router.Run(":9000"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
