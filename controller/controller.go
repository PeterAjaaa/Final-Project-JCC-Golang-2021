package controller

import (
	"strconv"

	"Final-Project-JCC-Golang-2021/entity"
	"Final-Project-JCC-Golang-2021/service"

	"github.com/gin-gonic/gin"
)

type GameController interface {
	FindAll() []entity.Game
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type gameController struct {
	service service.GameService
}

type UserController interface {
	FindAllUser() []entity.User
	SaveUser(ctx *gin.Context) error
	UpdateUser(ctx *gin.Context) error
	DeleteUser(ctx *gin.Context) error
}

type userController struct {
	userService service.UserService
}

type UrlController interface {
	FindAllUrl() []entity.GameURL
	SaveUrl(ctx *gin.Context) error
	UpdateUrl(ctx *gin.Context) error
	DeleteUrl(ctx *gin.Context) error
}

type urlController struct {
	urlService service.UrlService
}

type ReviewController interface {
	FindAllReview() []entity.Review
	SaveReview(ctx *gin.Context) error
	UpdateReview(ctx *gin.Context) error
	DeleteReview(ctx *gin.Context) error
}

type reviewController struct {
	reviewService service.ReviewService
}

type RatingController interface {
	FindAllRating() []entity.Rating
	SaveRating(ctx *gin.Context) error
	UpdateRating(ctx *gin.Context) error
	DeleteRating(ctx *gin.Context) error
}

type ratingController struct {
	ratingService service.RatingService
}

type StoreController interface {
	FindAllStore() []entity.Store
	SaveStore(ctx *gin.Context) error
	UpdateStore(ctx *gin.Context) error
	DeleteStore(ctx *gin.Context) error
}

type storeController struct {
	storeService service.StoreService
}

type AvailbilityController interface {
	FindAllAvailbility() []entity.Availbility
	SaveAvailbility(ctx *gin.Context) error
	UpdateAvailbility(ctx *gin.Context) error
	DeleteAvailbility(ctx *gin.Context) error
}

type availbilityController struct {
	availbilityService service.AvailbilityService
}

type PlatformController interface {
	FindAllPlatform() []entity.Platform
	SavePlatform(ctx *gin.Context) error
	UpdatePlatform(ctx *gin.Context) error
	DeletePlatform(ctx *gin.Context) error
}

type platformController struct {
	platformService service.PlatformService
}

// Game Functions

func New(service service.GameService) GameController {
	return &gameController{
		service: service,
	}
}

func (c *gameController) FindAll() []entity.Game {
	return c.service.FindAll()
}

func (c *gameController) Save(ctx *gin.Context) error {
	var game entity.Game
	if err := ctx.ShouldBindJSON(&game); err != nil {
		return err
	}
	c.service.Save(game)
	return nil
}

func (c *gameController) Update(ctx *gin.Context) error {
	var game entity.Game
	if err := ctx.ShouldBindJSON(&game); err != nil {
		return err
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	game.ID = id

	c.service.Update(game)
	return nil
}

func (c *gameController) Delete(ctx *gin.Context) error {
	var game entity.Game
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	game.ID = id
	c.service.Delete(game)
	return nil
}

// User Functions

func NewUser(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) FindAllUser() []entity.User {
	return c.userService.FindAllUser()
}

func (uc *userController) SaveUser(ctx *gin.Context) error {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return err
	}
	uc.userService.SaveUser(user)
	return nil
}

func (uc *userController) UpdateUser(ctx *gin.Context) error {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return err
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	user.ID = id

	uc.userService.UpdateUser(user)
	return nil
}

func (uc *userController) DeleteUser(ctx *gin.Context) error {
	var user entity.User
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	user.ID = id
	uc.userService.DeleteUser(user)
	return nil
}

// URL Functions

func NewUrl(urlService service.UrlService) UrlController {
	return &urlController{
		urlService: urlService,
	}
}

func (c *urlController) FindAllUrl() []entity.GameURL {
	return c.urlService.FindAllUrl()
}

func (uc *urlController) SaveUrl(ctx *gin.Context) error {
	var url entity.GameURL
	if err := ctx.ShouldBindJSON(&url); err != nil {
		return err
	}
	uc.urlService.SaveUrl(url)
	return nil
}

func (uc *urlController) UpdateUrl(ctx *gin.Context) error {
	var url entity.GameURL
	if err := ctx.ShouldBindJSON(&url); err != nil {
		return err
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	url.ID = id

	uc.urlService.UpdateUrl(url)
	return nil
}

func (uc *urlController) DeleteUrl(ctx *gin.Context) error {
	var url entity.GameURL
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	url.ID = id
	uc.urlService.DeleteUrl(url)
	return nil
}

// Review Functions

func NewReview(reviewService service.ReviewService) ReviewController {
	return &reviewController{
		reviewService: reviewService,
	}
}

func (c *reviewController) FindAllReview() []entity.Review {
	return c.reviewService.FindAllReview()
}

func (c *reviewController) SaveReview(ctx *gin.Context) error {
	var review entity.Review
	if err := ctx.ShouldBindJSON(&review); err != nil {
		return err
	}
	c.reviewService.SaveReview(review)
	return nil
}

func (c *reviewController) UpdateReview(ctx *gin.Context) error {
	var review entity.Review
	if err := ctx.ShouldBindJSON(&review); err != nil {
		return err
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	review.ID = id

	c.reviewService.UpdateReview(review)
	return nil
}

func (c *reviewController) DeleteReview(ctx *gin.Context) error {
	var review entity.Review
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	review.ID = id
	c.reviewService.DeleteReview(review)
	return nil
}

// Rating Functions

func NewRating(ratingService service.RatingService) RatingController {
	return &ratingController{
		ratingService: ratingService,
	}
}

func (c *ratingController) FindAllRating() []entity.Rating {
	return c.ratingService.FindAllRating()
}

func (c *ratingController) SaveRating(ctx *gin.Context) error {
	var rating entity.Rating
	if err := ctx.ShouldBindJSON(&rating); err != nil {
		return err
	}
	c.ratingService.SaveRating(rating)
	return nil
}

func (c *ratingController) UpdateRating(ctx *gin.Context) error {
	var rating entity.Rating
	if err := ctx.ShouldBindJSON(&rating); err != nil {
		return err
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	rating.ID = id

	c.ratingService.UpdateRating(rating)
	return nil
}

func (c *ratingController) DeleteRating(ctx *gin.Context) error {
	var rating entity.Rating
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	rating.ID = id
	c.ratingService.DeleteRating(rating)
	return nil
}

// Store Functions

func NewStore(storeService service.StoreService) StoreController {
	return &storeController{
		storeService: storeService,
	}
}

func (c *storeController) FindAllStore() []entity.Store {
	return c.storeService.FindAllStore()
}

func (c *storeController) SaveStore(ctx *gin.Context) error {
	var store entity.Store
	if err := ctx.ShouldBindJSON(&store); err != nil {
		return err
	}
	c.storeService.SaveStore(store)
	return nil
}

func (c *storeController) UpdateStore(ctx *gin.Context) error {
	var store entity.Store
	if err := ctx.ShouldBindJSON(&store); err != nil {
		return err
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	store.ID = id

	c.storeService.UpdateStore(store)
	return nil
}

func (c *storeController) DeleteStore(ctx *gin.Context) error {
	var store entity.Store
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	store.ID = id
	c.storeService.DeleteStore(store)
	return nil
}

// Availbility Functions

func NewAvailbility(availbilityService service.AvailbilityService) AvailbilityController {
	return &availbilityController{
		availbilityService: availbilityService,
	}
}

func (c *availbilityController) FindAllAvailbility() []entity.Availbility {
	return c.availbilityService.FindAllAvailbility()
}

func (c *availbilityController) SaveAvailbility(ctx *gin.Context) error {
	var availbility entity.Availbility
	if err := ctx.ShouldBindJSON(&availbility); err != nil {
		return err
	}
	c.availbilityService.SaveAvailbility(availbility)
	return nil
}

func (c *availbilityController) UpdateAvailbility(ctx *gin.Context) error {
	var availbility entity.Availbility
	if err := ctx.ShouldBindJSON(&availbility); err != nil {
		return err
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	availbility.ID = id

	c.availbilityService.UpdateAvailbility(availbility)
	return nil
}

func (c *availbilityController) DeleteAvailbility(ctx *gin.Context) error {
	var availbility entity.Availbility
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	availbility.ID = id
	c.availbilityService.DeleteAvailbility(availbility)
	return nil
}

// Platform Functions

func NewPlatform(platformService service.PlatformService) PlatformController {
	return &platformController{
		platformService: platformService,
	}
}

func (c *platformController) FindAllPlatform() []entity.Platform {
	return c.platformService.FindAllPlatform()
}

func (c *platformController) SavePlatform(ctx *gin.Context) error {
	var platform entity.Platform
	if err := ctx.ShouldBindJSON(&platform); err != nil {
		return err
	}
	c.platformService.SavePlatform(platform)
	return nil
}

func (c *platformController) UpdatePlatform(ctx *gin.Context) error {
	var platform entity.Platform
	if err := ctx.ShouldBindJSON(&platform); err != nil {
		return err
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	platform.ID = id

	c.platformService.UpdatePlatform(platform)
	return nil
}

func (c *platformController) DeletePlatform(ctx *gin.Context) error {
	var platform entity.Platform
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	platform.ID = id
	c.platformService.DeletePlatform(platform)
	return nil
}
