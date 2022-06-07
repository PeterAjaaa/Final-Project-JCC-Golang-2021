package service

import (
	"Final-Project-JCC-Golang-2021/entity"
	"Final-Project-JCC-Golang-2021/repository"
)

type GameService interface {
	Save(entity.Game) entity.Game
	Update(game entity.Game)
	Delete(game entity.Game)
	FindAll() []entity.Game
}

type gameService struct {
	gameRepository repository.GameRepository
}

type UserService interface {
	FindAllUser() []entity.User
	SaveUser(entity.User) entity.User
	UpdateUser(user entity.User)
	DeleteUser(user entity.User)
}

type userService struct {
	userRepository repository.UserRepository
}

type UrlService interface {
	FindAllUrl() []entity.GameURL
	SaveUrl(entity.GameURL) entity.GameURL
	UpdateUrl(url entity.GameURL)
	DeleteUrl(url entity.GameURL)
}

type urlService struct {
	urlRepository repository.UrlRepository
}

type ReviewService interface {
	SaveReview(entity.Review) entity.Review
	UpdateReview(review entity.Review)
	DeleteReview(review entity.Review)
	FindAllReview() []entity.Review
}

type reviewService struct {
	reviewRepository repository.ReviewRepository
}

type RatingService interface {
	SaveRating(entity.Rating) entity.Rating
	UpdateRating(rating entity.Rating)
	DeleteRating(rating entity.Rating)
	FindAllRating() []entity.Rating
}

type ratingService struct {
	ratingRepository repository.RatingRepository
}

type StoreService interface {
	FindAllStore() []entity.Store
	SaveStore(entity.Store) entity.Store
	UpdateStore(store entity.Store)
	DeleteStore(store entity.Store)
}

type storeService struct {
	storeRepository repository.StoreRepository
}

type AvailbilityService interface {
	SaveAvailbility(entity.Availbility) entity.Availbility
	UpdateAvailbility(availbility entity.Availbility)
	DeleteAvailbility(availbility entity.Availbility)
	FindAllAvailbility() []entity.Availbility
}

type availbilityService struct {
	availbilityRepository repository.AvailbilityRepository
}

type PlatformService interface {
	SavePlatform(entity.Platform) entity.Platform
	UpdatePlatform(platform entity.Platform)
	DeletePlatform(platform entity.Platform)
	FindAllPlatform() []entity.Platform
}

type platformService struct {
	platformRepository repository.PlatformRepository
}

// Game Functions

func New(repo repository.GameRepository) GameService {
	return &gameService{
		gameRepository: repo,
	}
}

func (service *gameService) Save(game entity.Game) entity.Game {
	service.gameRepository.Save(game)
	return game
}

func (service *gameService) Update(game entity.Game) {
	service.gameRepository.Update(game)
}

func (service *gameService) Delete(game entity.Game) {
	service.gameRepository.Delete(game)
}

func (service *gameService) FindAll() []entity.Game {
	return service.gameRepository.FindAll()
}

// User Functions

func NewUser(repo repository.UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}

func (service *userService) SaveUser(user entity.User) entity.User {
	service.userRepository.SaveUser(user)
	return user
}

func (service *userService) UpdateUser(user entity.User) {
	service.userRepository.UpdateUser(user)
}
func (service *userService) DeleteUser(user entity.User) {
	service.userRepository.DeleteUser(user)
}

func (service *userService) FindAllUser() []entity.User {
	return service.userRepository.FindAllUser()
}

// URL Functions

func NewUrl(repo repository.UrlRepository) UrlService {
	return &urlService{
		urlRepository: repo,
	}
}

func (service *urlService) SaveUrl(url entity.GameURL) entity.GameURL {
	service.urlRepository.SaveUrl(url)
	return url
}

func (service *urlService) UpdateUrl(url entity.GameURL) {
	service.urlRepository.UpdateUrl(url)
}

func (service *urlService) DeleteUrl(url entity.GameURL) {
	service.urlRepository.DeleteUrl(url)
}

func (service *urlService) FindAllUrl() []entity.GameURL {
	return service.urlRepository.FindAllUrl()
}

// Review Functions

func NewReview(repo repository.ReviewRepository) ReviewService {
	return &reviewService{
		reviewRepository: repo,
	}
}

func (service *reviewService) SaveReview(review entity.Review) entity.Review {
	service.reviewRepository.SaveReview(review)
	return review
}

func (service *reviewService) UpdateReview(review entity.Review) {
	service.reviewRepository.UpdateReview(review)
}

func (service *reviewService) DeleteReview(review entity.Review) {
	service.reviewRepository.DeleteReview(review)
}

func (service *reviewService) FindAllReview() []entity.Review {
	return service.reviewRepository.FindAllReview()
}

// Rating Functions

func NewRating(repo repository.RatingRepository) RatingService {
	return &ratingService{
		ratingRepository: repo,
	}
}

func (service *ratingService) SaveRating(rating entity.Rating) entity.Rating {
	service.ratingRepository.SaveRating(rating)
	return rating
}

func (service *ratingService) UpdateRating(rating entity.Rating) {
	service.ratingRepository.UpdateRating(rating)
}

func (service *ratingService) DeleteRating(rating entity.Rating) {
	service.ratingRepository.DeleteRating(rating)
}

func (service *ratingService) FindAllRating() []entity.Rating {
	return service.ratingRepository.FindAllRating()
}

// Store Functions

func NewStore(repo repository.StoreRepository) StoreService {
	return &storeService{
		storeRepository: repo,
	}
}

func (service *storeService) SaveStore(store entity.Store) entity.Store {
	service.storeRepository.SaveStore(store)
	return store
}

func (service *storeService) UpdateStore(store entity.Store) {
	service.storeRepository.UpdateStore(store)
}

func (service *storeService) DeleteStore(store entity.Store) {
	service.storeRepository.DeleteStore(store)
}

func (service *storeService) FindAllStore() []entity.Store {
	return service.storeRepository.FindAllStore()
}

// Availbility Functions

func NewAvailbility(repo repository.AvailbilityRepository) AvailbilityService {
	return &availbilityService{
		availbilityRepository: repo,
	}
}

func (service *availbilityService) SaveAvailbility(availbility entity.Availbility) entity.Availbility {
	service.availbilityRepository.SaveAvailbility(availbility)
	return availbility
}

func (service *availbilityService) UpdateAvailbility(availbility entity.Availbility) {
	service.availbilityRepository.UpdateAvailbility(availbility)
}

func (service *availbilityService) DeleteAvailbility(availbility entity.Availbility) {
	service.availbilityRepository.DeleteAvailbility(availbility)
}

func (service *availbilityService) FindAllAvailbility() []entity.Availbility {
	return service.availbilityRepository.FindAllAvailbility()
}

// Platform Functions

func NewPlatform(repo repository.PlatformRepository) PlatformService {
	return &platformService{
		platformRepository: repo,
	}
}

func (service *platformService) SavePlatform(platform entity.Platform) entity.Platform {
	service.platformRepository.SavePlatform(platform)
	return platform
}

func (service *platformService) UpdatePlatform(platform entity.Platform) {
	service.platformRepository.UpdatePlatform(platform)
}

func (service *platformService) DeletePlatform(platform entity.Platform) {
	service.platformRepository.DeletePlatform(platform)
}

func (service *platformService) FindAllPlatform() []entity.Platform {
	return service.platformRepository.FindAllPlatform()
}
