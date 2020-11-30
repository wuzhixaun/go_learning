package pipe_filter

import "testing"

func TestNewStraightPipeline(t *testing.T) {
	splitFilter := NewSplitFilter(",")
	toIntFilter := NewToIntFilter()
	sumFilter := NewSumFilter()

	pipeline := NewStraightPipeline("p1", splitFilter, toIntFilter, sumFilter)

	process, _ := pipeline.Process("1,2,3,4,5")

	t.Log(process)
}
