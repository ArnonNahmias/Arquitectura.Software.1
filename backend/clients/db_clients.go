package clients

import (
	"backend/dao"
	"fmt"
	"time"
)

const (
	tableNameUsers         = "users"
	tableNameCourses       = "courses"
	tableNameSubscriptions = "subscriptions"
)

func main() {
	db := map[string][]interface{}{
		tableNameUsers: {
			dao.User{
				ID:           1,
				Email:        "arnonahmias13@gmail.com",
				PasswordHash: "5f4dcc3b5aa765d61d8327deb882cf99",
				Type:         "admin",
				CreationDate: time.Now().UTC(),
				LastUpdated:  time.Now().UTC(),
			},

			dao.User{
				ID:           2,
				Email:        "juanperez@gmail.com",
				PasswordHash: "5f4dcc3b5aa765d61d8327deb882cf99",
				Type:         "normal",
				CreationDate: time.Now().UTC(),
			},
		},
	}

	// ... Additional user data and operations
	users := db[tableNameUsers].([]dao.User)
	fmt.Println("User 1:", users[0])
	fmt.Println("User 2:", users[1]) // This will fail
}
