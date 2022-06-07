# Documentation For Games Data & Review API

## | Super Users account credentials for Heroku deployment:

**USERNAME=admin**
**PASSWORD=admin**

**USERNAME2=trusteduser**
**PASSWORD2=secretfortrusteduser**

**USERNAME3=esrbfolks**
**PASSWORD3=passforesrbfolks**

***Fields marked by "(SU)" means Super User credentials are needed to access the associated pages (or operations)***

***If you're running the API locally, set the SU credentials (username & password) using .env file on your own machine***

## | Games
GET : https://calm-atoll-29984.herokuapp.com/games (Get all Games)

POST : https://calm-atoll-29984.herokuapp.com/games (Add new Games) (SU)

PUT : https://calm-atoll-29984.herokuapp.com/games/:id (Update Games by ID) (SU)

DELETE : https://calm-atoll-29984.herokuapp.com/games/:id (Delete Games by ID) (SU)

### | Games Data
    ID          int64
    Title       string
    ReleaseYear int
    Description string

## | User
GET : https://calm-atoll-29984.herokuapp.com/user (Get all User) (SU)

POST : https://calm-atoll-29984.herokuapp.com/user (Add new User)

PUT : https://calm-atoll-29984.herokuapp.com/user/:id (Update User by ID)

DELETE : https://calm-atoll-29984.herokuapp.com/user/:id (Delete User by ID) (SU)

### | User Data
    ID        int64
	Name      string
	Email     string
	Password  string

## | URL
GET : https://calm-atoll-29984.herokuapp.com/URL (Get all URL)

POST : https://calm-atoll-29984.herokuapp.com/URL (Add new URL) (SU)

PUT : https://calm-atoll-29984.herokuapp.com/URL/:id (Update URL by ID) (SU)

DELETE : https://calm-atoll-29984.herokuapp.com/URL/:id (Delete URL by ID) (SU)

### | URL Data
    ID        int64
	CoverURL  string
    Game      Game

## | Review
GET : https://calm-atoll-29984.herokuapp.com/review (Get all Review)

POST : https://calm-atoll-29984.herokuapp.com/review (Add new Review) (SU)

PUT : https://calm-atoll-29984.herokuapp.com/review/:id (Update Review by ID) (SU)

DELETE : https://calm-atoll-29984.herokuapp.com/review/:id (Delete Review by ID) (SU)

### | Review Data
    ID         int64
	UserReview string
    Game       Game

## | Rating
GET : https://calm-atoll-29984.herokuapp.com/rating (Get all Rating)

POST : https://calm-atoll-29984.herokuapp.com/rating (Add new Rating)

PUT : https://calm-atoll-29984.herokuapp.com/rating/:id (Update Rating by ID)

DELETE : https://calm-atoll-29984.herokuapp.com/rating/:id (Delete Rating by ID) (SU)

### | Rating Data
    ID         int64
	UserRating int
	ESRBRating string
	User       User
	Game       Game

## | Store
GET : https://calm-atoll-29984.herokuapp.com/store (Get all Store)

POST : https://calm-atoll-29984.herokuapp.com/store (Add new Store) (SU)

PUT : https://calm-atoll-29984.herokuapp.com/store/:id (Update Store by ID) (SU)

DELETE : https://calm-atoll-29984.herokuapp.com/store/:id (Delete Store by ID) (SU)

### | Store Data
    ID         int64
	SteamCheck int
	Game       Game

## | Availbility
GET : https://calm-atoll-29984.herokuapp.com/availbility (Get all Availbility)

POST : https://calm-atoll-29984.herokuapp.com/availbility (Add new Availbility) (SU)

PUT : https://calm-atoll-29984.herokuapp.com/availbility/:id (Update Availbility by ID) (SU)

DELETE : https://calm-atoll-29984.herokuapp.com/availbility/:id (Delete Availbility by ID) (SU)

### | Availbility Data
    ID        int64
	Stock     string
	Game      Game

## | Platform
GET : https://calm-atoll-29984.herokuapp.com/platform (Get all Platform)

POST : https://calm-atoll-29984.herokuapp.com/platform (Add new Platform) (SU)

PUT : https://calm-atoll-29984.herokuapp.com/platform/:id (Update Platform by ID) (SU)

DELETE : https://calm-atoll-29984.herokuapp.com/platform/:id (Delete Platform by ID) (SU)

### | Platform Data
    ID            int64
	CrossPlatform int
	Game          Game