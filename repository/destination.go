package repository

import (
	"database/sql"
	"fmt"

	"github.com/train-do/Golang-Generics/model"
	"github.com/train-do/Golang-Generics/utils"
)

type RepoDestination struct {
	Db *sql.DB
}

func NewRepoDestination(db *sql.DB) *RepoDestination {
	return &RepoDestination{db}
}
func (r *RepoDestination) FindAll(qp model.QueryParams) ([]model.Destinations, int, error) {
	fmt.Printf("%+v------\n", qp)
	var args []any
	page, sort, search := utils.GenerateQuery(qp, &args)
	subQuery := fmt.Sprintf(`
	with "total_order" as(
	select d.id as "destination_id", d.name as "destination_name", count(o.id) as "total_order"
	from "Order" o
	left join "DestinationSchedule" ds ON ds.id = o.destination_schedule_id
	left join "Destination" d ON d.id = ds.destination_id
	group by d.id),
	"rating" as(
	select d.id as "destination_id", d.name as "destination_name", round(avg(o.rating),1) as "average_rating"
	from "Order" o
	left join "DestinationSchedule" ds ON ds.id = o.destination_schedule_id
	left join "Destination" d ON d.id = ds.destination_id
	group by d.id),
	"total_items" as (
	select count(*) as "total_items"
	from "Destination" d
	join "DestinationSchedule" ds on ds.destination_id = d.id
	join "Schedule" s on s.id = ds.schedule_id
	left join "total_order" t on d.id = t.destination_id
	left join "rating" r on d.id = r.destination_id %s)`, search)
	query := fmt.Sprintf(`%s select ds.id , d."name" , d.description , d.image_url  , s."date" , ds.price , t.total_order, r.average_rating, ti.total_items
	from "Destination" d
	join "DestinationSchedule" ds on ds.destination_id = d.id
	join "Schedule" s on s.id = ds.schedule_id
	left join "total_order" t on d.id = t.destination_id
	left join "rating" r on d.id = r.destination_id
	left join "total_items" ti on true %s %s %s`, subQuery, search, sort, page)
	// fmt.Println(query, args)
	var destinations []model.Destinations
	var totalItem int
	rows, err := r.Db.Query(query, args...)
	if err != nil {
		fmt.Println("Error Query : ", err)
		return []model.Destinations{}, 0, err
	}
	for rows.Next() {
		var destination model.Destinations
		rows.Scan(&destination.Id, &destination.Name, &destination.Description, &destination.ImageUrl, &destination.Date, &destination.Price, &destination.TotalBooking, &destination.Rating, &totalItem)
		destinations = append(destinations, destination)
	}
	return destinations, totalItem, nil
}
func (r *RepoDestination) FindById(destination *model.Destination, id int) error {
	query := `with "rating" as(
	select d.id as "destination_id", d.name as "destination_name", round(avg(o.rating),1) as "average_rating", count(*) 
	from "Order" o
	left join "DestinationSchedule" ds ON ds.id = o.destination_schedule_id
	left join "Destination" d ON d.id = ds.destination_id
	group by d.id)
	select ds.id , d."name" , d.description , d.image_url  , s."date" , ds.price , r.average_rating, r.count, g.id , g.description , g.image_url 
	from "Destination" d 
	join "DestinationSchedule" ds on ds.destination_id = d.id 
	join "Schedule" s on s.id = ds.schedule_id
	left join "rating" r on d.id = r.destination_id
	left join "Gallery" g on g.destination_id = d.id
	where ds.id = $1;`
	rows, err := r.Db.Query(query, id)
	if err != nil {
		fmt.Println("Error Query : ", err)
		return err
	}
	for rows.Next() {
		var gallery model.Gallery
		rows.Scan(&destination.DestinationScheduleId, &destination.Name, &destination.Description, &destination.ImageUrl, &destination.Date, &destination.Price, &destination.Rating, &destination.TotalReview, &gallery.Id, &gallery.Description, &gallery.ImageUrl)
		destination.Gallery = append(destination.Gallery, gallery)
	}
	return nil
}
