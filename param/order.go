package param

import "time"

type Order struct {
	CustomerName *string    `json:"customer_name"`
	Items        []Item     `json:"items,omitempty"`
	OrderedAt    *time.Time `json:"orderedAt,omitempty"`
}
