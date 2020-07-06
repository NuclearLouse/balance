package models

import "time"

// Client ...
type Client struct {
	Type      int
	Name      string
	User      string
	Markup    float64
	Status    bool
	Comment   string
	CreatedAt time.Time
}

func (c Client) String() string {
	switch c.Type {
	case 1:
		return "поставщик"
	case 2:
		return "покупатель"
	case 3:
		return "перевозчик"
	}
	return "неизвестный тип клиента"
}

