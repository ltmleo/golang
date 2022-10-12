package main

import (
	"fmt"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

func main(){
	product := Product{
		ID: "1",
		Name: "Product 1",
		Price: 10.98,
	}
	err := SaveProduct(product)
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.POST("/products/", createProduct)
	e.Logger.Fatal
	fmt.Println(product.Name)
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil) // GO ROUTINE
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World!!"))
}

func CreateProduct(c echo.Context) error {
	product := Product{}
	if err := c.Bind(&product); err != nil {
		return err
	}
	return c.Json()
}

func SaveProduct(product Product) error {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO products(id,name,price) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
}