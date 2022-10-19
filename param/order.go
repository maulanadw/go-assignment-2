package param

import "time"

type Order struct {
	ID           int        `json:"id,omitempty"`
	CustomerName *string    `json:"customer_name"`
	Items        []Item     `json:"items,omitempty"`
	OrderedAt    *time.Time `json:"ordered_at,omitempty"`
}
