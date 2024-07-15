package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var err error

type User struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

func main() {
	// Connect to the MySQL database
	db, err = sqlx.Connect("mysql", "root:mysql#pass@tcp(localhost:3306)/test")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	r := gin.Default()

	// Define routes that interact with the database
	r.GET("/user-test", GetUsers)
	r.POST("/user-test", CreateUser)

	r.Run(":8800")
}

func GetUsers(c *gin.Context) {
	var users []User
	err := db.Select(&users, "SELECT id, name, description FROM user_test")
	if err != nil {
		c.JSON(500, gin.H{"error": "Error retrieving users"})
		return
	}
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	_, err := db.Exec("INSERT INTO user_test (name, description) VALUES (?, ?)", user.Name, user.Description)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating user"})
		return
	}
	c.JSON(200, user)
}

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// func main() {
// 	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 	db, err := gorm.Open("mysql", `root:mysql#pass@tcp(localhost:3306)/test?charset=utf8&parseTime=True`)

// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	db.AutoMigrate(&Book{})

// 	handler := newHandler(db)

// 	r := gin.New()

// 	r.GET("/books", handler.listBooksHandler)
// 	r.POST("/books", handler.createBookHandler)
// 	r.DELETE("/books/:id", handler.deleteBookHandler)

// 	r.Run()
// }

// type Handler struct {
// 	db *gorm.DB
// }

// func newHandler(db *gorm.DB) *Handler {
// 	return &Handler{db}
// }

// type Book struct {
// 	ID     string `json:"id"`
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// }

// func (h *Handler) listBooksHandler(c *gin.Context) {
// 	var books []Book

// 	if result := h.db.Find(&books); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, &books)
// }

// func (h *Handler) createBookHandler(c *gin.Context) {
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

// func (h *Handler) deleteBookHandler(c *gin.Context) {
// 	id := c.Param("id")

// 	if result := h.db.Delete(&Book{}, id); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": result.Error.Error(),
// 		})
// 		return
// 	}

// 	c.Status(http.StatusNoContent)
// }

// import (
// 	"log"
// 	"net/http"

// 	"github.com/labstack/echo"
// 	"gorm.io/gorm"
// )

// func main() {
// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})

// 	h := CustomerHandler{}
// 	h.Initialize()

// 	e.GET("/customers", h.GetAllCustomer)
// 	e.POST("/customers", h.SaveCustomer)
// 	e.GET("/customers/:id", h.GetCustomer)
// 	e.PUT("/customers/:id", h.UpdateCustomer)
// 	e.DELETE("/customers/:id", h.DeleteCustomer)

// 	e.Logger.Fatal(e.Start(":8800"))
// }

// type CustomerHandler struct {
// 	DB *gorm.DB
// }

// func (h *CustomerHandler) Initialize() {
// 	db, err := gorm.Open("mysql", `root:mysql#pass@tcp(localhost:3306)/test?charset=utf8&parseTime=True`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	db.AutoMigrate(&Customer{})

// 	h.DB = db
// }

// type Customer struct {
// 	Id        uint   `gorm:"primary_key" json:"id"`
// 	FirstName string `json:"firstName"`
// 	LastName  string `json:"lastName"`
// 	Age       int    `json:"age"`
// 	Email     string `json:"email"`
// }

// func (h *CustomerHandler) GetAllCustomer(c echo.Context) error {
// 	customers := []Customer{}

// 	h.DB.Find(&customers)

// 	return c.JSON(http.StatusOK, customers)
// }

// func (h *CustomerHandler) GetCustomer(c echo.Context) error {
// 	id := c.Param("id")
// 	customer := Customer{}

// 	if err := h.DB.Find(&customer, id).Error; err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	return c.JSON(http.StatusOK, customer)
// }

// func (h *CustomerHandler) SaveCustomer(c echo.Context) error {
// 	customer := Customer{}

// 	if err := c.Bind(&customer); err != nil {
// 		return c.NoContent(http.StatusBadRequest)
// 	}

// 	if err := h.DB.Save(&customer).Error; err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}

// 	return c.JSON(http.StatusOK, customer)
// }

// func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
// 	id := c.Param("id")
// 	customer := Customer{}

// 	if err := h.DB.Find(&customer, id).Error; err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	if err := c.Bind(&customer); err != nil {
// 		return c.NoContent(http.StatusBadRequest)
// 	}

// 	if err := h.DB.Save(&customer).Error; err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}

// 	return c.JSON(http.StatusOK, customer)
// }

// func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
// 	id := c.Param("id")
// 	customer := Customer{}

// 	if err := h.DB.Find(&customer, id).Error; err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	if err := h.DB.Delete(&customer).Error; err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}

// 	return c.NoContent(http.StatusNoContent)
// }

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
