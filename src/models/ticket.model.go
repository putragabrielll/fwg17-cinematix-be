package models

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type ticket struct{
	Title string `db:"title" json:"title"`
	Rating string `db:"name" json:"rating"`
	Date time.Time `db:"date" json:"date"`
	Time time.Time `db:"time" json:"time"`
	SeatCount int `db:"seatCount" json:"seatCount"`
	Total int `db:"total" json:"total"`
	SeatCode pq.StringArray `db:"seatCode" json:"seatCode"`
}



func GetTicket(c *gin.Context, orderId int, userId int) (ticket, error) {
	sql := `
	SELECT
	"m"."title",
	"r"."name",
	"d"."date",
	"at"."time",
	"o"."seatCount",
	"o"."total",
	array_agg(DISTINCT "od"."seatCode") "seatCode"
	FROM "order" "o"
	JOIN "orderDetail" "od" ON ("od"."orderId" = "o"."id")
	JOIN "movies" "m" ON ("m"."id" = "o"."moviesId")
	JOIN "rating" "r" ON ("r"."id" = "m"."ratingId")
	JOIN "movieCinema" "mc" ON ("mc"."moviesId" = "m"."id")
	JOIN "cinema" "c" ON ("c"."id" = "mc"."cinemaId")
	JOIN "moviesTime" "mt" ON ("mt"."cinemaId" = "c"."id")
	JOIN "airingTimeDate" "atd" ON ("atd"."id" = "mt"."airingTimeDateId")
	JOIN "airingTime" "at" ON ("at"."id" = "atd"."airingTimeId")
	JOIN "date" "d" ON ("d"."id" = "atd"."dateId")
	WHERE "o"."id" = $1 AND "o"."usersId" = $2
	GROUP BY "m"."title", "r"."name", "d"."date", "at"."time", "o"."seatCount", "o"."total"
	`
	data := ticket{}
	err := db.Get(&data, sql, orderId, userId)
	if err != nil {
		return data, err
	}
	
	return data, err
}