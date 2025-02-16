package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/st-ember/mockecommerce/internal/db"
	"github.com/st-ember/mockecommerce/internal/init_db/storage"
)

func main() {
	fmt.Println("Initializing DB Schema and Content")

	// Read SQL file from disk
	sqlFile, err := os.ReadFile("mock_ecommerce_init.sql")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connecting to Postgres")

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	db.InitDB(os.Getenv("CONN_STR"))
	defer db.CloseDB()

	_, err = db.DB.Exec(string(sqlFile))
    if err != nil {
        log.Fatal(err)
    }

	err = storage.StoreInitCountries()
	if err != nil {
		log.Fatal("Error Initializing Countries", err)
	}

	err = storage.StoreInitCustomers()
	if err != nil {
		log.Fatal("Error Initializing Customers", err)
	}

	err = storage.StoreInitMerchants()
	if err != nil {
		log.Fatal("Error Initializing Merchants", err)
	}

	err = storage.StoreInitProductCategories()
	if err != nil {
		log.Fatal("Error Initializing Product Categories", err)
	}

	err = storage.StoreInitProducts()
	if err != nil {
		log.Fatal("Error Initializing Products", err)
	}
}
