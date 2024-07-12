package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	/*e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})*/
	h := UserTestHandler{}
	h.Initialize()

	e.GET("/user-test", h.GetAllUserTest)
	e.POST("/user-test", h.SaveUserTest)
	e.GET("user-test/:id", h.GetUserTest)
	e.PUT("/user-test/:id", h.UpdateUserTest)
	e.DELETE("/user-test/:id", h.DeleteUserTest)

	e.Logger.Fatal(e.Start(":8800"))
}

type UserTestHandler struct {
	DB *gorm.DB
}

// ให้เชื่อมต่อฐานข้อมูลเมื่อ Initialize
func (h *UserTestHandler) Initialize() {
	db, err := gorm.Open("mysql", "root:mysql#pass@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	h.DB = db
}

type User struct {
	Id          uint   `gorm:"primary_key" json:"id"`
	Name        string `json:"firstName"`
	Description string `json:"lastName"`
}

func (h *UserTestHandler) GetAllUserTest(c echo.Context) error {
	users := []User{}

	h.DB.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func (h *UserTestHandler) GetUserTest(c echo.Context) error {
	id := c.Param("id")
	user := User{}

	if err := h.DB.Find(&user, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserTestHandler) SaveUserTest(c echo.Context) error {
	user := User{}

	if err := c.Bind(&user); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&user); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserTestHandler) UpdateUserTest(c echo.Context) error {
	id := c.Param("id")
	user := User{}

	if err := h.DB.Find(&user, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := c.Bind(&user).Error; err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&user).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserTestHandler) DeleteUserTest(c echo.Context) error {
	id := c.Param("id")
	user := User{}

	if err := h.DB.Find(&user, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}

// import (
// 	"database/sql"
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	fmt.Println("Hello, World!")
// 	val_instance := gin.Default()
// 	val_instance.GET("/", http_Handler)
// 	val_instance.GET("/db-test", getAllTestModels)
// 	val_instance.Run()
// }

// func setupDB() (*sql.DB, error) {
// 	connectionString := "root:mysql#pass@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
// 	database, _ := sql.Open("mysql", connectionString)
// 	return database, nil
// }

// type Test_Detail struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// func getAllTestModels(c *gin.Context) {
// 	database, _ := setupDB()
// 	defer database.Close()
// 	rows, _ := database.Query("SELECT * FROM user_test")
// 	defer rows.Close()
// 	var data []Test_Detail
// 	for rows.Next() {
// 		var name string
// 		var description string
// 		var id int
// 		rows.Scan(&id, &name, &description)
// 		data = append(data, Test_Detail{ID: id, Name: name, Description: description})
// 	}
// 	c.JSON(http.StatusOK, data)
// }

// func http_Handler(http_instance *gin.Context) {
// 	http_instance.JSON(200, gin.H{
// 		"message": "TESTING",
// 	})
// }

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	_ "github.com/go-sql-driver/mysql"
// )

// const (
// 	username = "root"
// 	password = "mysql#pass"
// 	hostname = "127.0.0.1:3306"
// 	dbname   = "test"
// )

// func dsn(dbName string) string {
// 	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
// }

// func main() {
// 	db, err := sql.Open("mysql", dsn(dbname))
// 	if err != nil {
// 		log.Printf("Error %s when opening DB", err)
// 		return
// 	}
// 	defer db.Close()

// 	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancelfunc()

// 	err = db.PingContext(ctx)
// 	if err != nil {
// 		log.Printf("Errors %s pinging DB", err)
// 		return
// 	}

// 	log.Printf("Connected to DB successfully\n")

// 	// _, err = db.ExecContext(ctx, `CREATE TABLE Product (
// 	// 	id INT NOT NULL AUTO_INCREMENT,
// 	// 	product_code VARCHAR(45) NOT NULL,
// 	// 	product_name VARCHAR(45) NOT NULL,
// 	// 	quantity INT NOT NULL,
// 	// 	PRIMARY KEY (id));
// 	//   `)
// 	// if err != nil {
// 	// 	log.Printf("Error %s when creating Table\n", err)
// 	// 	return
// 	// }
// 	// log.Printf("Create table successfully\n")

// 	return
// }

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// func main() {
// 	var err error
// 	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	db.AutoMigrate(&Book{})

// 	r := gin.New()

// 	r.GET("/books", listBooksHandler)
// 	r.POST("/books", createBookHandler)
// 	r.DELETE("/books/:id", deleteBookHandler)

// 	r.Run()
// }

// type Book struct {
// 	ID     string `json:"id"`
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// }

// func listBooksHandler(c *gin.Context) {
// 	var books []Book

// 	if result := db.Find(&books); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, &books)
// }
// func createBookHandler(c *gin.Context) {
// 	var book Book

// 	if err := c.ShouldBindJSON(&book); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	if result := db.Create(&book); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, &book)
// }
// func deleteBookHandler(c *gin.Context) {
// 	id := c.Param("id")

// 	if result := db.Delete(&Book{}, id); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	c.Status(http.StatusNoContent)
// }

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Hello World!")
// }

// func main() {
// 	r := gin.New()

// 	r.GET("/users", listUsersHandler)
// 	r.POST("/users", createUserHandler)
// 	r.DELETE("/users/:id", deleteUserHandler)

// 	r.Run()
// }

// func main() {
// 	r := gin.Default()
// 	gin.SetMode(gin.ReleaseMode)

// 	// Connect to your database
// 	db, err := gorm.Open("mysql", "root:mysql#pass@tcp(localhost:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic("Failed to connect to database")
// 	}

// 	defer db.Close()

// 	// Your API routes and handlers
// 	r.Run(":8080")
// }

// type User struct {
// 	ID          string `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// var users = []User{
// 	{ID: "1", Name: "test no.1", Description: "King"},
// 	{ID: "2", Name: "test no.2", Description: "Queen"},
// 	{ID: "3", Name: "test no.3", Description: "Jack"},
// }

// func listUsersHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, books)
// }

// func createUserHandler(c *gin.Context) {
// 	var user User

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	users = append(users, user)

// 	c.JSON(http.StatusCreated, user)
// }

// func deleteUserHandler(c *gin.Context) {
// 	id := c.Param("id")

// 	for i, a := range users {
// 		if a.ID == id {
// 			users = append(users[:i], users[i+1:]...)
// 			break
// 		}
// 	}

//		c.Status(http.StatusNoContent)
//	}

// var db *sql.DB

// func main() {
// 	db = connectDB()
// 	defer db.Close()
// 	router := mux.NewRouter()
// 	router.HandleFunc("/users", getUsers).Methods("GET")
// 	router.HandleFunc("/users/{id}", getUser).Methods("GET")
// 	router.HandleFunc("/users", createUser).Methods("POST")
// 	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
// 	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
// 	log.Fatal(http.ListenAndServe(":8000", router))
// }
// func connectDB() *sql.DB {
// 	// Replace "username", "password", "dbname" with your database credentials
// 	connectionString := "root:mysql#pass@tcp(localhost:3306)/test"
// 	db, err := sql.Open("mysql", connectionString)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db
// }
// func getUsers(w http.ResponseWriter, r *http.Request) {
// 	rows, err := db.Query("SELECT * FROM users")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()
// 	// Iterate through the rows and create a JSON response
// 	// You can use a JSON library like "encoding/json" to marshal the data
// }
// func getUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	userID := vars["id"]
// 	row := db.QueryRow("SELECT * FROM users WHERE id = ?", userID)
// 	// Parse the row data and create a JSON response
// }
// func createUser(w http.ResponseWriter, r *http.Request) {
// 	// Parse the request body and extract user data
// 	// Insert the user data into the "users" table
// 	// Handle validation and error cases
// }
// func updateUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	userID := vars["id"]
// 	// Parse the request body and extract user data
// 	// Update the user data in the "users" table based on the user ID
// 	// Handle validation and error cases
// }
// func deleteUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	userID := vars["id"]
// 	_, err := db.Exec("DELETE FROM users WHERE id = ?", userID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	// Send a success response
// }

// func main() {
// 	r := gin.New()
// 	r.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Hello World!",
// 		})
// 	})
// 	r.GET("/users", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, users)
// 	})

// 	r.POST("/users", func(c *gin.Context) {
// 		var user User

// 		if err := c.ShouldBindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}
// 		users = append(users, user)
// 		c.JSON(http.StatusCreated, user)
// 	})

// 	r.DELETE("/users/:id", func(c *gin.Context) {
// 		id := c.Param("id")

// 		for i, a := range users {
// 			if a.ID == id {
// 				users = append(users[:i], users[i+1:]...)
// 				break
// 			}
// 		}
// 		c.Status(http.StatusNoContent)
// 	})

// 	r.Run(":8080")
// }
