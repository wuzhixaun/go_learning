package pipe_filter

import (
	"errors"
	"strings"
)

type SplitFilter struct {
	delimiter string // 分隔符
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {

	str, ok := data.(string) // 检验数据格式/类型，是否可以处理
	if !ok {                 // SplitFilterFormateError
		return nil, errors.New("SplitFilterFormateError")
	}

	split := strings.Split(str, sf.delimiter)
	return split, nil

}
