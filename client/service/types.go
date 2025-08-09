package service

type PizzaOrder struct {
	Flavor   string  `json:"flavor"`
	Size     string  `json:"size"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Address  string  `json:"address"`
	UserName string  `json:"user_name"`
}
