package pipe_filter

import "errors"

type SumFilter struct {
}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
	ints, ok := data.([]int)

	if !ok {
		return nil, errors.New("data. error")
	}

	sum := 0
	for _, v := range ints {
		sum += v
	}

	return sum, nil
}
