package param

type Item struct {
	ItemCode    *string `json:"item_code"`
	Description *string `json:"description"`
	Quantity    *int    `json:"quantity"`
}
