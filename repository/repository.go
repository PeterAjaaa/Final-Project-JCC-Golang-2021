package repository

import (
	"Final-Project-JCC-Golang-2021/entity"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GameRepository interface {
	Save(game entity.Game)
	Update(game entity.Game)
	Delete(game entity.Game)
	FindAll() []entity.Game
	CloseDB()
}

type UserRepository interface {
	FindAllUser() []entity.User
	SaveUser(user entity.User)
	UpdateUser(user entity.User)
	DeleteUser(user entity.User)
	CloseDB()
}

type UrlRepository interface {
	FindAllUrl() []entity.GameURL
	SaveUrl(url entity.GameURL)
	UpdateUrl(url entity.GameURL)
	DeleteUrl(url entity.GameURL)
	CloseDB()
}

type ReviewRepository interface {
	SaveReview(review entity.Review)
	UpdateReview(review entity.Review)
	DeleteReview(review entity.Review)
	FindAllReview() []entity.Review
	CloseDB()
}

type RatingRepository interface {
	SaveRating(rating entity.Rating)
	UpdateRating(rating entity.Rating)
	DeleteRating(rating entity.Rating)
	FindAllRating() []entity.Rating
	CloseDB()
}

type StoreRepository interface {
	FindAllStore() []entity.Store
	SaveStore(store entity.Store)
	UpdateStore(store entity.Store)
	DeleteStore(store entity.Store)
	CloseDB()
}

type AvailbilityRepository interface {
	SaveAvailbility(availbility entity.Availbility)
	UpdateAvailbility(availbility entity.Availbility)
	DeleteAvailbility(availbility entity.Availbility)
	FindAllAvailbility() []entity.Availbility
	CloseDB()
}

type PlatformRepository interface {
	SavePlatform(platform entity.Platform)
	UpdatePlatform(platform entity.Platform)
	DeletePlatform(platform entity.Platform)
	FindAllPlatform() []entity.Platform
	CloseDB()
}
type database struct {
	connection *gorm.DB
}

func (db *database) CloseDB() {
	sqlDB, err := db.connection.DB()
	sqlDB.Close()
	if err != nil {
		panic("Failed to close MySQL database!")
	}
}

func NewGameRepository() GameRepository {
	godotenv.Load()
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to open MySQL database!")
	}
	db.AutoMigrate(&entity.Game{})
	return &database{
		connection: db,
	}
}

func (db *database) Save(game entity.Game) {
	db.connection.Create(&game)
}

func (db *database) Update(game entity.Game) {
	db.connection.Save(&game)
}

func (db *database) Delete(game entity.Game) {
	db.connection.Delete(&game)
}

func (db *database) FindAll() []entity.Game {
	var games []entity.Game
	db.connection.Set("gorm:auto_preload", true).Find(&games)
	return games
}

// User Functions

func NewUserRepository() UserRepository {
	godotenv.Load()
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to open MySQL database!")
	}
	db.AutoMigrate(&entity.User{})
	return &database{
		connection: db,
	}
}

func (db *database) SaveUser(user entity.User) {
	db.connection.Create(&user)
}

func (db *database) UpdateUser(user entity.User) {
	db.connection.Save(&user)
}

func (db *database) DeleteUser(user entity.User) {
	db.connection.Delete(&user)
}

func (db *database) FindAllUser() []entity.User {
	var user []entity.User
	db.connection.Set("gorm:auto_preload", true).Find(&user)
	return user
}

// URL Functions

func NewUrlRepository() UrlRepository {
	godotenv.Load()
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to open MySQL database!")
	}
	db.AutoMigrate(&entity.GameURL{}, &entity.Game{})
	return &database{
		connection: db,
	}
}

func (db *database) SaveUrl(url entity.GameURL) {
	db.connection.Create(&url)
}

func (db *database) UpdateUrl(url entity.GameURL) {
	db.connection.Save(&url)
}

func (db *database) DeleteUrl(url entity.GameURL) {
	db.connection.Delete(&url)
}

func (db *database) FindAllUrl() []entity.GameURL {
	var url []entity.GameURL
	db.connection.Set("gorm:auto_preload", true).Find(&url)
	return url
}

// Review Functions

func NewReview() ReviewRepository {
	godotenv.Load()
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to open MySQL database!")
	}
	db.AutoMigrate(&entity.Review{}, &entity.Game{})
	return &database{
		connection: db,
	}
}

func (db *database) SaveReview(review entity.Review) {
	db.connection.Create(&review)
}

func (db *database) UpdateReview(review entity.Review) {
	db.connection.Save(&review)
}

func (db *database) DeleteReview(review entity.Review) {
	db.connection.Delete(&review)
}

func (db *database) FindAllReview() []entity.Review {
	var review []entity.Review
	db.connection.Set("gorm:auto_preload", true).Find(&review)
	return review
}

// Rating Functions

func NewRating() RatingRepository {
	godotenv.Load()
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to open MySQL database!")
	}
	db.AutoMigrate(&entity.Rating{}, &entity.User{}, &entity.Game{})
	return &database{
		connection: db,
	}
}

func (db *database) SaveRating(review entity.Rating) {
	db.connection.Create(&review)
}

func (db *database) UpdateRating(review entity.Rating) {
	db.connection.Save(&review)
}

func (db *database) DeleteRating(review entity.Rating) {
	db.connection.Delete(&review)
}

func (db *database) FindAllRating() []entity.Rating {
	var rating []entity.Rating
	db.connection.Set("gorm:auto_preload", true).Find(&rating)
	return rating
}

// Store Functions
func NewStore() StoreRepository {
	godotenv.Load()
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to open MySQL database!")
	}
	db.AutoMigrate(&entity.Store{}, &entity.Game{})
	return &database{
		connection: db,
	}
}

func (db *database) SaveStore(review entity.Store) {
	db.connection.Create(&review)
}

func (db *database) UpdateStore(review entity.Store) {
	db.connection.Save(&review)
}

func (db *database) DeleteStore(review entity.Store) {
	db.connection.Delete(&review)
}

func (db *database) FindAllStore() []entity.Store {
	var store []entity.Store
	db.connection.Set("gorm:auto_preload", true).Find(&store)
	return store
}

// Availbility Functions

func NewAvailbility() AvailbilityRepository {
	godotenv.Load()
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to open MySQL database!")
	}
	db.AutoMigrate(&entity.Availbility{})
	return &database{
		connection: db,
	}
}

func (db *database) SaveAvailbility(availbility entity.Availbility) {
	db.connection.Create(&availbility)
}

func (db *database) UpdateAvailbility(availbility entity.Availbility) {
	db.connection.Save(&availbility)
}

func (db *database) DeleteAvailbility(availbility entity.Availbility) {
	db.connection.Delete(&availbility)
}

func (db *database) FindAllAvailbility() []entity.Availbility {
	var availbility []entity.Availbility
	db.connection.Set("gorm:auto_preload", true).Find(&availbility)
	return availbility
}

// Platform Functions

func NewPlatform() PlatformRepository {
	godotenv.Load()
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to open MySQL database!")
	}
	db.AutoMigrate(&entity.Platform{})
	return &database{
		connection: db,
	}
}

func (db *database) SavePlatform(platform entity.Platform) {
	db.connection.Create(&platform)
}

func (db *database) UpdatePlatform(platform entity.Platform) {
	db.connection.Save(&platform)
}

func (db *database) DeletePlatform(platform entity.Platform) {
	db.connection.Delete(&platform)
}

func (db *database) FindAllPlatform() []entity.Platform {
	var platform []entity.Platform
	db.connection.Set("gorm:auto_preload", true).Find(&platform)
	return platform
}
