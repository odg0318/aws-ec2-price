package price

type Instance struct {
	Region string  `json:"region"`
	Type   string  `json:"type"`
	Price  float64 `json:"price"`
}
