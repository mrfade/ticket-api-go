package main

import (
	"github.com/mrfade/ticket-api-go/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
}

func main() {
	initializers.SyncDatabase()
}
