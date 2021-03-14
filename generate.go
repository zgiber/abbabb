package abbabb

import (
	"math/rand"
	"strings"
)

// returns the gapped copy as string as it was in the test
func GenerateExample(l int) string {
	complete := GenerateCompleteExample(l)
	gapped := &strings.Builder{}
	for i := 0; i < l; i++ {
		if rand.Intn(3) == 0 {
			gapped.WriteString("?")
			continue
		}
		switch complete[i] {
		case 1:
			gapped.WriteString("a")
		case -1:
			gapped.WriteString("b")
		}
	}

	return gapped.String()
}

func GenerateCompleteExample(lenght int) []int {
	s := make([]int, lenght)
	for i := 0; i < lenght; i++ {
		if i < 2 {
			s[i] = random()
			continue
		}

		if sum(s[i-2:i+1]) >= 2 {
			s[i] = -1
			continue
		}

		if sum(s[i-2:i+1]) <= -2 {
			s[i] = 1
			continue
		}
		s[i] = random()
	}
	return s
}

func sum(n []int) int {
	if len(n) == 0 {
		return 0
	}

	sum := n[0]
	for i := 1; i < len(n); i++ {
		sum += n[i]
	}
	return sum
}

func random() int {
	n := rand.Intn(2)
	if n == 0 {
		return -1
	}
	return n
}
