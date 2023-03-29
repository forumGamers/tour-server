package main

import (
	cfg "github.com/forumGamers/tour-service/config"
	r "github.com/forumGamers/tour-service/routes"
)

func main() {
	cfg.Connection()

	r.Routes()
}