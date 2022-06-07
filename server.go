package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"Final-Project-JCC-Golang-2021/controller"
	"Final-Project-JCC-Golang-2021/middleware"
	"Final-Project-JCC-Golang-2021/repository"
	"Final-Project-JCC-Golang-2021/service"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	gameRepository repository.GameRepository = repository.NewGameRepository()
	gameService    service.GameService       = service.New(gameRepository)
	GameController controller.GameController = controller.New(gameService)

	userRepository repository.UserRepository = repository.NewUserRepository()
	userService    service.UserService       = service.NewUser(userRepository)
	UserController controller.UserController = controller.NewUser(userService)

	urlRepository repository.UrlRepository = repository.NewUrlRepository()
	urlService    service.UrlService       = service.NewUrl(urlRepository)
	UrlController controller.UrlController = controller.NewUrl(urlService)

	reviewRepository repository.ReviewRepository = repository.NewReview()
	reviewService    service.ReviewService       = service.NewReview(reviewRepository)
	ReviewController controller.ReviewController = controller.NewReview(reviewService)

	ratingRepository repository.RatingRepository = repository.NewRating()
	ratingService    service.RatingService       = service.NewRating(ratingRepository)
	RatingController controller.RatingController = controller.NewRating(ratingService)

	storeRepository repository.StoreRepository = repository.NewStore()
	storeService    service.StoreService       = service.NewStore(storeRepository)
	StoreController controller.StoreController = controller.NewStore(storeService)

	availbilityRepository repository.AvailbilityRepository = repository.NewAvailbility()
	availbilityService    service.AvailbilityService       = service.NewAvailbility(availbilityRepository)
	AvailbilityController controller.AvailbilityController = controller.NewAvailbility(availbilityService)

	platformRepository repository.PlatformRepository = repository.NewPlatform()
	platformService    service.PlatformService       = service.NewPlatform(platformRepository)
	PlatformController controller.PlatformController = controller.NewPlatform(platformService)
)

func main() {
	defer gameRepository.CloseDB()
	defer userRepository.CloseDB()
	defer urlRepository.CloseDB()
	defer reviewRepository.CloseDB()
	defer ratingRepository.CloseDB()
	defer storeRepository.CloseDB()
	defer availbilityRepository.CloseDB()
	defer platformRepository.CloseDB()

	middleware.SetupLogOutput()

	server := gin.New()
	server.Use(middleware.Logger(), gin.Recovery(), gindump.Dump())

	superUsers := server.Group("/", middleware.SuperUsersAuth())

	port := os.Getenv("PORT")

	// Start Of Games function

	server.GET("/games", func(c *gin.Context) {
		c.JSON(http.StatusOK, GameController.FindAll())
	})

	superUsers.POST("/games", func(c *gin.Context) {
		if err := GameController.Save(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	superUsers.PUT("/games/:id/", func(c *gin.Context) {
		if err := GameController.Update(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	superUsers.DELETE("/games/:id/", func(c *gin.Context) {
		if err := GameController.Delete(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	// Start of User Functions

	superUsers.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, UserController.FindAllUser())
	})

	server.POST("/user", func(c *gin.Context) {
		if err := UserController.SaveUser(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	server.PUT("/user/:id/", func(c *gin.Context) {
		if err := UserController.UpdateUser(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	superUsers.DELETE("/user/:id/", func(c *gin.Context) {
		if err := UserController.DeleteUser(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	// Start of URL Functions

	server.GET("/URL", func(c *gin.Context) {
		c.JSON(http.StatusOK, UrlController.FindAllUrl())
	})

	superUsers.POST("/URL", func(c *gin.Context) {
		if err := UrlController.SaveUrl(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	superUsers.PUT("/URL/:id/", func(c *gin.Context) {
		if err := UrlController.UpdateUrl(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	superUsers.DELETE("/URL/:id/", func(c *gin.Context) {
		if err := UrlController.DeleteUrl(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	// Start of Review Functions

	server.GET("/review", func(c *gin.Context) {
		c.JSON(http.StatusOK, ReviewController.FindAllReview())
	})

	superUsers.POST("/review", func(c *gin.Context) {
		if err := ReviewController.SaveReview(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	superUsers.PUT("/review/:id/", func(c *gin.Context) {
		if err := ReviewController.UpdateReview(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	superUsers.DELETE("/review/:id/", func(c *gin.Context) {
		if err := ReviewController.DeleteReview(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	// Start of Rating Functions

	server.GET("/rating", func(c *gin.Context) {
		c.JSON(http.StatusOK, RatingController.FindAllRating())
	})

	server.POST("/rating", func(c *gin.Context) {
		if err := RatingController.SaveRating(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	server.PUT("/rating/:id/", func(c *gin.Context) {
		if err := RatingController.UpdateRating(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	superUsers.DELETE("/rating/:id/", func(c *gin.Context) {
		if err := RatingController.DeleteRating(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	// Start of Store Functions

	server.GET("/store", func(c *gin.Context) {
		c.JSON(http.StatusOK, StoreController.FindAllStore())
	})

	superUsers.POST("/store", func(c *gin.Context) {
		if err := StoreController.SaveStore(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	superUsers.PUT("/store/:id/", func(c *gin.Context) {
		if err := StoreController.UpdateStore(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	superUsers.DELETE("/store/:id/", func(c *gin.Context) {
		if err := StoreController.DeleteStore(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	// Start of Availbility Functions

	server.GET("/availbility", func(c *gin.Context) {
		c.JSON(http.StatusOK, AvailbilityController.FindAllAvailbility())
	})

	superUsers.POST("/availbility", func(c *gin.Context) {
		if err := AvailbilityController.SaveAvailbility(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	superUsers.PUT("/availbility/:id/", func(c *gin.Context) {
		if err := AvailbilityController.UpdateAvailbility(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	superUsers.DELETE("/availbility/:id/", func(c *gin.Context) {
		if err := AvailbilityController.DeleteAvailbility(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	// Start of Platform Functions

	server.GET("/platform", func(c *gin.Context) {
		c.JSON(http.StatusOK, PlatformController.FindAllPlatform())
	})

	superUsers.POST("/platform", func(c *gin.Context) {
		if err := PlatformController.SavePlatform(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	superUsers.PUT("/platform/:id/", func(c *gin.Context) {
		if err := PlatformController.UpdatePlatform(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	superUsers.DELETE("/platform/:id/", func(c *gin.Context) {
		if err := PlatformController.DeletePlatform(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"Message": "Input Valid!"})
	})

	fmt.Println("Server is now running!")
	log.Fatal(server.Run(":" + port))
}
