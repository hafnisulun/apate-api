package models

type Residence struct {
	Base
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

type ResidencesResponseBody struct {
	Data []Residence `json:"data"`
}

type ResidenceResponseBody struct {
	Data Residence `json:"data"`
}

type FindResidencesQuery struct {
	Page    int `form:"page"`
	PerPage int `form:"per_page"`
}
