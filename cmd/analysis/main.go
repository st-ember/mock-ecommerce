package analysis

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/st-ember/mockecommerce/internal/db"
)

func main() {
	fmt.Println("Connecting to Postgres")

	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	db.InitDB(os.Getenv("CONN_STR"))
	defer db.CloseDB()
}
