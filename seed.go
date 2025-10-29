package main

import (
	"context"
	"log"
	"sql_c/pkg/repository"

	"github.com/bxcodec/faker/v3"
)

func (r *Repo) seed() {
	for i := 0; i < 10; i++ {

		_, err := r.db.CreateUser(context.Background(), repository.CreateUserParams{
			Email:        faker.Email(),
			PasswordHash: faker.Password(),
			PhoneNumber:  faker.E164PhoneNumber(),
		})
		if err != nil {
			log.Fatalf("error: %v", err.Error())
		}
	}
}
