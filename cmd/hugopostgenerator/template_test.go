package hugopostgenerator

import (
	"testing"
	"time"
)

func TestCreateTemplate(t *testing.T) {
	d1 := Details{
		Date: time.Now(),
		Venue: Venue{
			Name:    "Concert Venue Name",
			Address: "Some address 43, 4536MK Amsterdam",
		},
		Ticket: Ticket{
			Description: "Reserve Ticket",
			Link:        "https://sellmyticket.money",
			Price:       "78E",
		},
	}
	d2 := Details{
		Date: time.Time{},
		Venue: Venue{
			Name:    "Another concert Venue",
			Address: "Another address 89 7824OP Utrecht",
		},
		Ticket: Ticket{
			Description: "Buy ticket",
			Link:        "https://anotherlinkforticket.nl",
			Price:       "97E",
		},
	}
	td := &templateData{
		Date:          time.Now(),
		Description:   "Test Description",
		FeaturedImage: "test_image.jpeg",
		Tags: []string{
			"music",
			"concert",
		},
		Title: "Test concert title",
		Year:  "2098",
		Details: []Details{
			d1,
			d2,
		},
		Content: "Some long content in markdown",
	}

	t.Run("create template", func(t *testing.T) {
		createTemplate(td)
	})
}
