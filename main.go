package main

import (
	"fmt"
)

type Rental struct {
	ID       int
	Bookings []Booking
}

type Booking struct {
	StartDate string
	EndDate   string
}

type requestedDates struct {
	StartDate string
	EndDate   string
}

type availableRentalsForDates struct {
	StartDate string
	EndDate   string
	RentalIDs []int
}

func main() {
	inputs := []requestedDates{
		{
			StartDate: "2023-07-24",
			EndDate:   "2023-07-27",
		},
		{
			StartDate: "2023-07-22",
			EndDate:   "2023-07-27",
		},
		{
			StartDate: "2023-07-20",
			EndDate:   "2023-07-23",
		},
		{
			StartDate: "2023-08-08",
			EndDate:   "2023-08-10",
		},
		{
			StartDate: "2023-08-16",
			EndDate:   "2023-08-20",
		},
		{
			StartDate: "2023-08-08",
			EndDate:   "2023-08-12",
		},
	}

	rentals := []Rental{
		{
			ID: 1,
			Bookings: []Booking{
				{
					StartDate: "2023-07-23",
					EndDate:   "2023-07-26",
				},
				{
					StartDate: "2023-07-28",
					EndDate:   "2023-08-02",
				},
				{
					StartDate: "2023-08-05",
					EndDate:   "2023-08-07",
				},
				{
					StartDate: "2023-08-11",
					EndDate:   "2023-08-17",
				},
			},
		},
		{
			ID: 2,
			Bookings: []Booking{
				{
					StartDate: "2023-07-25",
					EndDate:   "2023-07-27",
				},
				{
					StartDate: "2023-07-29",
					EndDate:   "2023-08-03",
				},
				{
					StartDate: "2023-08-06",
					EndDate:   "2023-08-08",
				},
				{
					StartDate: "2023-08-12",
					EndDate:   "2023-08-14",
				},
			},
		},
	}

	// write a function that based on the rentals bookings and requested dates
	// returns a slice of requested dates with rental ids that are available for those dates
	// bookings are considered blocked days that are currently unavailable during that time
	// the bookings can share the same start or end date
	// a requested date range that passes through unavailable dates, should not be available

	available := fetchAvailabilityForRequestedDates(inputs, rentals)
	fmt.Println(available)
}

func fetchAvailabilityForRequestedDates(inputs []requestedDates, rentals []Rental) []availableRentalsForDates {
	availableDates := []availableRentalsForDates{}

	// TODO

	return availableDates
}
