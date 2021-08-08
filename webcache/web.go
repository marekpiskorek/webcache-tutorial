package webcache

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func GetDBString() string {
	Host := os.Getenv("DBHOST")
	Port := os.Getenv("DBPORT")
	User := os.Getenv("DBUSER")
	Name := os.Getenv("DBNAME")
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		Host, Port, User, Name)
}

func CachedWebpageHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := pgx.Connect(context.Background(), GetDBString())
	checkErr(err)
	defer conn.Close(context.Background())

	var text string
	created := time.Now()

	err = conn.QueryRow(context.Background(), "SELECT cached_content from cached_webpage").Scan(&text)
	checkErr(err)

	if err != nil {
		noCacheText := "No cache found"
		newCacheText := fmt.Sprintf("This is a cached webpage from %s", created.Format(time.ANSIC))
		_, err := conn.Exec(context.Background(), "INSERT INTO cached_webpage(cached_content) values($1)", newCacheText)
		checkErr(err)

		fmt.Fprintf(w, noCacheText)
	} else {
		fmt.Fprintf(w, text)
	}
}
