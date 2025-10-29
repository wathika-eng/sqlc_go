package seed

import (
	"context"
	"log"
	"sql_c/pkg/repository"
	"strings"

	"github.com/bxcodec/faker/v3"
)

func Seeder(r *repository.Queries) {
	for i := 0; i < 2; i++ {

		_, err := r.CreateUser(context.Background(), repository.CreateUserParams{
			Email:        strings.ToLower(faker.Email()),
			PasswordHash: faker.Password(),
			PhoneNumber:  faker.E164PhoneNumber(),
		})
		if err != nil {
			log.Fatalf("error: %v", err.Error())
		}
		log.Println("done")
	}
}
