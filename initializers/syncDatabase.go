package initializers

import "github.com/mrfade/ticket-api-go/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Cast{})
	DB.AutoMigrate(&models.City{})
	DB.AutoMigrate(&models.Crew{})
	DB.AutoMigrate(&models.Genre{})
	DB.AutoMigrate(&models.MovieGenre{})
	DB.AutoMigrate(&models.Movie{})
	DB.AutoMigrate(&models.MovieSession{})
	DB.AutoMigrate(&models.MovieSessionSeat{})
	DB.AutoMigrate(&models.Person{})
	DB.AutoMigrate(&models.Place{})
	DB.AutoMigrate(&models.Theater{})
	DB.AutoMigrate(&models.TheaterPrice{})
	DB.AutoMigrate(&models.TheaterSeat{})
	DB.AutoMigrate(&models.Ticket{})
	DB.AutoMigrate(&models.User{})
}
