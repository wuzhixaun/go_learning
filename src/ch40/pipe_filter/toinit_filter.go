package pipe_filter

import (
	"errors"
	"strconv"
)

type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (tif *ToIntFilter) Process(data Request) (Response, error) {
	strings, ok := data.([]string)

	if !ok {
		return nil, errors.New("NewToIntFilter errot")
	}

	ret := []int{}

	for _, part := range strings {
		atoi, err := strconv.Atoi(part)

		if err != nil {
			return nil, errors.New("strconv.Atoi error")
		}

		ret = append(ret, atoi)
	}

	return ret, nil

}
