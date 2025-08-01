package meta

type Query struct {
	OrderBy        string
	OrderDirection string
	Status         string
	Limit          int
	Offset         int
}

func Parse(metadata *Params) (*Query, error) {

	limit := metadata.PerPage
	offset := (metadata.Page - 1) * limit

	q := Query{
		OrderBy:        metadata.OrderBy,
		OrderDirection: metadata.OrderType,
		Limit:          limit,
		Offset:         offset,
	}

	return &q, nil
}
