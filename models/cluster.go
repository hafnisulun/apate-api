package models

type Cluster struct {
	Base
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

type ClustersResponseBody struct {
	Data []Cluster `json:"data"`
}

type ClusterResponseBody struct {
	Data Cluster `json:"data"`
}

type FindClustersQuery struct {
	Page    int `form:"page"`
	PerPage int `form:"per_page"`
}
