package server

import (
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"sync"
)

type Campsite struct {
	mu       sync.Mutex
	bookings []Booking
}

func NewCampsite() *Campsite {
	return &Campsite{}
}

func (c *Campsite) Create(booking Booking) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	booking.UUID = uuid.New().String()
	c.bookings = append(c.bookings, booking)
	return booking.UUID, nil
}

func (c *Campsite) Read(uuid string) (Booking, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	idx := slices.IndexFunc(c.bookings, func(b Booking) bool { return b.UUID == uuid })
	if idx == -1 {
		return Booking{}, ErrBookingNotFound
	}
	return c.bookings[idx], nil
}

// Booking START:types
type Booking struct {
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	FullName  string `json:"fullName"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Active    bool   `json:"active"`
}

//END:types

var ErrBookingNotFound = fmt.Errorf("booking not found")
