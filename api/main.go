package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NewUser struct {
	Id        string `json:"id" binding:"-"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Avatar    string `josn:"avatar" binding:"-"`
	CreatedAt string `josn:"created_at" binding:"-"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Id        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Avatar    string    `gorm:"default:null"`
	CreatedAt string    `gorm:"type:timestamptz;default:now()"`
	Comments  []Comment `gorm:"foreignKey:UserId;references:Id"`
}

type Comment struct {
	Id        uint  `gorm:"type:int autoIncrement primaryKey"`
	ParentId  *uint `gorm:"type:int;default:null"`
	Text      string
	Upvotes   uint
	CreatedAt string `gorm:"type:timestamptz;default:now()"`
	UserId    string `gorm:"type:uuid"`
}

func main() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = false
	corsConfig.AllowOrigins = []string{"http://localhost:5173", "https://openion.netlify.app/"}
	corsConfig.AllowCredentials = true
	corsConfig.ExposeHeaders = append(corsConfig.ExposeHeaders, "Set-Cookie")
	router.Use(cors.New(corsConfig))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DBURL")), &gorm.Config{TranslateError: true})
	jwt_key := []byte(os.Getenv("JWTKEY"))

	if err != nil {
		log.Fatal("failed to connect to db", err)
	}

	router.POST("/login", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx := context.Background()
		json.Email = strings.ToLower(json.Email)
		user, err := gorm.G[User](db).Where("email = ?", json.Email).First(ctx)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password)) == nil {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"iss": "openion",
				"sub": user.Id,
				"iat": jwt.NewNumericDate(time.Now()),
			})
			s, err := token.SignedString(jwt_key)
			if err != nil {
				log.Println("failed to sign jwt: %w", err)
				c.Status(http.StatusInternalServerError)
				return
			}
			c.SetCookie("jwt", s, 30*24*3600, "/", "localhost", false, true)
			c.JSON(http.StatusOK, gin.H{
				"UserId":   user.Id,
				"Name":     user.Name,
				"Email":    user.Email,
				"Avatar":   user.Avatar,
				"Password": "",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		}
	})

	router.GET("/getuser", JWTAuth(jwt_key, db), func(c *gin.Context) {
		ctx := context.Background()
		user_id, _ := c.Get("user_id")
		user, err := gorm.G[User](db).Select("id, name, email, avatar").Where("id = ?", user_id).First(ctx)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		c.JSON(http.StatusOK, user)

	})

	router.POST("/newuser", func(c *gin.Context) {
		var json NewUser
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		json.Email = strings.ToLower(json.Email)
		hashpw, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("failed to hash password: %w", err)
			c.Status(http.StatusInternalServerError)
		}
		user := User{Name: json.Name, Email: json.Email, Password: string(hashpw), Avatar: json.Avatar}

		ctx := context.Background()
		result := gorm.WithResult()
		err = gorm.G[User](db, result).Create(ctx, &user)

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusConflict, gin.H{"error": "The email address already exists"})
			return
		}

		user.Password = ""

		c.JSON(http.StatusOK, user)
	})

	router.GET("/comments", func(c *gin.Context) {
		pageStr := c.DefaultQuery("page", "1")
		pageSizeStr := c.DefaultQuery("pageSize", "10")
		parentIdStr := c.Query("parent_id")

		page, err := strconv.Atoi(pageStr)
		if err != nil || page <= 0 {
			page = 1
		}
		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil || pageSize <= 0 {
			pageSize = 10
		}
		offset := (page - 1) * pageSize

		query := gorm.G[map[string]any](db).Table("comments").Select("comments.id, parent_id, text, upvotes, comments.created_at, users.name AS user_name, users.avatar AS user_avatar").Joins(clause.JoinTarget{Association: "inner join users on comments.user_id = users.id"}, nil).Limit(pageSize).Offset(offset).Order("created_at ASC")

		counterQuery := gorm.G[map[string]any](db).Select("COUNT(r.id) AS replies").Table("comments c").Joins(clause.JoinTarget{Association: "left join comments r on r.parent_id = c.id"}, nil).Group("c.id").Order("c.created_at").Limit(pageSize)

		if offset != 0 {
			query = query.Offset(offset)
			counterQuery = counterQuery.Offset(offset)
		}

		if parentIdStr != "" {
			parentId, err := strconv.Atoi(parentIdStr)
			if err != nil || parentId == 0 {
				query = query.Where("parent_id is null")
				counterQuery = counterQuery.Where("c.parent_id is null")
			} else {
				query = query.Where("parent_id = ?", parentId)
				counterQuery = counterQuery.Where("c.parent_id = ?", parentId)
			}
		}

		ctx := context.Background()
		result, err := query.Find(ctx)
		if len(result) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"page":      nil,
				"page_size": pageSize,
				"data":      result,
			})
			return
		}

		ids := make([]any, 0, pageSize)
		for _, v := range result {
			ids = append(ids, v["id"])
		}

		if len(result) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"page":      nil,
				"page_size": pageSize,
				"data":      result,
			})
			return
		}

		counter, err := counterQuery.Find(ctx)

		for i := range len(result) {
			result[i]["replies"] = counter[i]["replies"]
		}

		c.JSON(http.StatusOK, gin.H{
			"page":      page,
			"page_size": pageSize,
			"data":      result,
		})
	})

	router.GET("/logout", func(c *gin.Context) {
		c.SetCookie("jwt", "", 0, "/", "localhost", false, true)
	})

	router.Run()
}

func JWTAuth(jwt_key []byte, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, err := c.Cookie("jwt")
		if err == nil {
			token, err := jwt.ParseWithClaims(val, &jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
				return jwt_key, nil
			}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

			if err != nil {
				c.Status(http.StatusUnauthorized)
				c.Abort()
				return
			}
			if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
				c.Set("user_id", claims.Subject)
				c.Next()
			}
		}
		c.Status(http.StatusOK)
		c.Abort()
	}
}
