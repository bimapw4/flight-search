package meta

import (
	"net/url"
	"strconv"
)

// SearchParams struct for holding search query parameters
type Params struct {
	Pagination
}

// NewSearchParams initializes search params from the request
func NewParams(queries map[string]string) Params {
	queryValues := url.Values{}
	for key, value := range queries {
		queryValues.Add(key, value)
	}

	return Params{
		Pagination: newPagination(queryValues),
	}
}

type Pagination struct {
	OrderType        string `json:"order_type"`
	OrderBy          string `json:"order_by"`
	Page             int    `json:"page"`
	PerPage          int    `json:"per_page"`
	TotalItems       int    `json:"total_items"`
	TotalTakeHomePay int    `json:"total_take_home_pay,omitempty"`
	SearchBy         string `json:"search_by"`
	Search           string `json:"search"`
}

// NewPagination initializes pagination info
func newPagination(query url.Values) Pagination {

	p := Pagination{
		PerPage:   10,
		Page:      1,
		OrderBy:   "created_at",
		OrderType: "asc",
	}

	if query.Get("order_by") != "" {
		p.OrderBy = query.Get("order_by")
	}

	if query.Get("order_type") != "" {
		p.OrderType = query.Get("order_type")
	}

	if query.Get("search_by") != "" {
		p.SearchBy = query.Get("search_by")
	}

	if query.Get("search") != "" {
		p.Search = query.Get("search")
	}

	pps := query.Get("per_page")
	if v, err := strconv.Atoi(pps); err == nil {
		if v <= 0 {
			v = 10
		}

		p.PerPage = v
	}

	ps := query.Get("page")
	if v, err := strconv.Atoi(ps); err == nil {
		if v < 1 {
			v = 1
		}

		p.Page = v
	}

	return p
}
