package Error

import (
	"log"
	"os"
)

func Handle_error(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
