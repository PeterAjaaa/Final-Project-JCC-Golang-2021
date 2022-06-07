package entity

import "time"

type User struct {
	ID        int64     `json:"user_id" gorm:"primary_key" type:"int(64)"`
	Name      string    `json:"name" binding:"required" gorm:"type:varchar(255)"`
	Email     string    `json:"email" binding:"required, email" gorm:"type:varchar(255)"`
	Password  string    `json:"pass" binding:"required" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Game struct {
	ID          int64     `json:"game_id" gorm:"primary_key" type:"int(64)"`
	Title       string    `json:"game_title" binding:"required" gorm:"type:varchar(255)"`
	ReleaseYear int       `json:"release_year" binding:"required"`
	Description string    `json:"description" binding:"required" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GameURL struct {
	ID        int64     `json:"url_id" gorm:"primary_key" type:"int(64)"`
	CoverURL  string    `json:"cover_url" binding:"required, url" gorm:"type:varchar(255)"`
	Game      Game      `gorm:"foreignKey:CoverURL"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Review struct {
	ID         int64     `json:"review_id" gorm:"primary_key" type:"int(64)"`
	UserReview string    `json:"user_review" binding:"required" gorm:"type:varchar(255)"`
	Game       Game      `gorm:"foreignKey:UserReview"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Rating struct {
	ID         int64     `json:"rating_id" gorm:"primary_key" type:"int(64)"`
	UserRating int       `json:"user_rating" binding:"required"`
	ESRBRating string    `json:"esrb_rating" binding:"required" gorm:"type:varchar(255)"`
	User       User      `gorm:"foreignKey:UserRating"`
	Game       Game      `gorm:"foreignKey:ESRBRating"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Store struct {
	ID         int64     `json:"store_id" gorm:"primary_key" type:"int(64)"`
	SteamCheck int       `json:"steam_check" binding:"required"`
	Game       Game      `gorm:"foreignKey:SteamCheck"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Availbility struct {
	ID        int64     `json:"avaibility_id" gorm:"primary_key" type:"int(64)"`
	Stock     string    `json:"stock" binding:"required" gorm:"type:varchar(255)"`
	Game      Game      `gorm:"foreignKey:Stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Platform struct {
	ID            int64     `json:"platform_id" gorm:"primary_key" type:"int(64)"`
	CrossPlatform int       `json:"cross_platform" binding:"required"`
	Game          Game      `gorm:"foreignKey:CrossPlatform"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
