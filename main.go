package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

var jwtSecret = []byte("your_secret_key")

func main() {
	e := echo.New()

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//add endpoints for the static files
	e.Static("/static", "static")
	e.Static("/views", "views")

	e.GET("/users", func(c echo.Context) error {
		rows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		defer rows.Close()

		var users []map[string]interface{}
		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			users = append(users, map[string]interface{}{
				"id":   id,
				"name": name,
			})
		}
		return c.JSON(http.StatusOK, users)
	})

	e.GET("/", func(c echo.Context) error {
		return c.File("views/index.html")
	})

	e.POST("/login", login)
	e.POST("/register", register)

	r := e.Group("/restricted")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: jwtSecret,
	}))

	r.GET("/character-sheets", getAllCharacterSheets)
	r.GET("/character-sheets/:id", getCharacterSheet)
	r.POST("/character-sheets", createCharacterSheet)
	r.PUT("/character-sheets/:id", updateCharacterSheet)
	r.DELETE("/character-sheets/:id", deleteCharacterSheet)

	//add endpoints for templates
	//add endpoints for the character sheets api
	//make an endpoint for getting all character sheets for a user
	//make an endpoint for getting a single character sheet
	//make an endpoint for creating a character sheet
	//make an endpoint for updating a character sheet
	//make an endpoint for deleting a character sheet
	//make an endpoint for logging in
	//make an endpoint for logging out
	//make an endpoint for registering a new user
	//make an endpoint for updating a user
	//make an endpoint for deleting a user

	e.Logger.Fatal(e.Start(":8080"))
}

func login(c echo.Context) error {
	// Implement login logic here
	// Check username and password in database
	// If the credentials are valid, generate JWT token
	// If the credentials are invalid, return an error

	// On successful login, generate and return JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func register(c echo.Context) error {
	// Implement user registration logic here
	return c.JSON(http.StatusOK, "User registered")
}

func getAllCharacterSheets(c echo.Context) error {
	// Implement logic to get all character sheets for a user
	return c.JSON(http.StatusOK, "All character sheets")
}

func getCharacterSheet(c echo.Context) error {
	// Implement logic to get a single character sheet
	return c.JSON(http.StatusOK, "Single character sheet")
}

func createCharacterSheet(c echo.Context) error {
	// Implement logic to create a character sheet
	return c.JSON(http.StatusOK, "Character sheet created")
}

func updateCharacterSheet(c echo.Context) error {
	// Implement logic to update a character sheet
	return c.JSON(http.StatusOK, "Character sheet updated")
}

func deleteCharacterSheet(c echo.Context) error {
	// Implement logic to delete a character sheet
	return c.JSON(http.StatusOK, "Character sheet deleted")
}

//add a function that renders a pdf based on a user's info

//add a function for dice rolls
//roll(dice)
//maybe add roll bonuses in a different place or in this function and return the changed roll
