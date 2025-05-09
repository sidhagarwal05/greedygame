package model

type Campaign struct {
	ID    string `json:"cid"`
	Image string `json:"img"`
	CTA   string `json:"cta"`
}

type DeliveryRequest struct {
	App     string
	Country string
	OS      string
}
