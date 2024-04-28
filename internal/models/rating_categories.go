package models

type Rating struct {
	CategoryID int
	Score      float64
}

type CategoryRating struct {
	TicketID      int32
	Ratings       map[int32]float64
	CategoryNames map[int32]string
}
