package model

import "time"

type Destinations struct {
	Id           int
	Name         string
	Description  string
	ImageUrl     string
	Date         time.Time
	Price        int
	TotalBooking int
	Rating       float32
}
type Destination struct {
	DestinationScheduleId int
	Name                  string
	Description           string
	ImageUrl              string
	Date                  time.Time
	Price                 int
	Rating                float32
	TotalReview           int
	Gallery               []Gallery
}
type QueryParams struct {
	Page        int
	SortDate    bool
	SortPrice   string
	SortName    bool
	SearchPlace string
	SearchDate  string
	SearchPrice int
}
type Gallery struct {
	Id          int
	Description string
	ImageUrl    string
}
