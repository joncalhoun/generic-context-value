package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joncalhoun/generic-context-value/context"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	return http.ListenAndServe(":3000", PokemonMiddleware(UserMiddleware(Handler)))
}

type Pokemon struct {
	Name  string
	Color string
}

type User struct {
	Email string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := context.Value[User](ctx)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	pokemon, err := context.Value[Pokemon](ctx)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, "Hello %s! You have a %s %s pokemon.\n", user.Email, pokemon.Color, pokemon.Name)
}

func PokemonMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// pretend to lookup a pokemon related to this request
		pokemon := Pokemon{
			Name:  "Charizard",
			Color: "Aquamarine",
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, pokemon)
		next(w, r.WithContext(ctx))
	}
}

func UserMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// pretend to lookup the current user
		user := User{
			Email: "jon@calhoun.io",
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, user)
		next(w, r.WithContext(ctx))
	}
}
