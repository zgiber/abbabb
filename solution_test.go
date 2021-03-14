package abbabb

import "testing"

func TestSolution(t *testing.T) {
	// rough tests for eyeballing if we're doing okay
	//
	s := Solution("abab???????ababab")
	t.Log(s)
	s = Solution("aabb???bb???ababab")
	t.Log(s)

	gapped := GenerateExample(1000)
	s = Solution(gapped)
	t.Log(gapped)
	t.Log(s)
}
