package abbabb

import (
	"math/rand"
)

var candidates = []int{-1, 1}

func Solution(slice string) string {
	weights := map[rune]int{
		'a': -1,
		'b': 1,
		'?': 0,
	}
	l := len(slice)
	weighted := make([]int, l)
	for i, s := range slice {
		weighted[i] = weights[s]
	}

	// helper keeping i within boundaries
	limit := func(n int) int {
		if n < 0 {
			return 0
		}
		if n > l {
			return l
		}
		return n
	}

	for i := 0; i < l; i++ {
		if weighted[i] != 0 {
			continue
		}

		switch sum(weighted[limit(i-1):limit(i+2)]) {
		case 2:
			// if sum of middle triplet is 2, new item is -1
			weighted[i] = -1
			continue
		case -2:
			// if sum of middle triplet is -2 new item is 1
			weighted[i] = 1
			continue
		}

		switch sum(weighted[limit(i-2):limit(i+1)]) {
		case 2:
			// if sum of left triplet is -2 then new item is 1
			weighted[i] = -1
			continue
		case -2:
			// if sum of left triplet is 2 then new item is -1
			weighted[i] = 1
			continue
		}

		switch sum(weighted[limit(i):limit(i+3)]) {
		case 2:
			// if sum of right triplet is -2 then new item is 1
			weighted[i] = -1
			continue
		case -2:
			// if sum of right triplet is 2 then new item is -1
			weighted[i] = 1
			continue
		}

		// otherwise new item is anything that moves middle triplet sum towards 0
		switch sum(weighted[limit(i-1):limit(i+2)]) {
		case 1:
			// if sum of middle triplet is 2, new item is -1
			weighted[i] = -1
			continue
		case -1:
			// if sum of middle triplet is -2 new item is 1
			weighted[i] = 1
			continue
		}

		// if we got here we can assign a random value
		weighted[i] = candidates[rand.Intn(2)]
	}

	res := make([]uint8, l)
	for i, v := range weighted {
		if v == -1 {
			res[i] = 'a'
			continue
		}

		if v == 1 {
			res[i] = 'b'
		}
	}

	return string(res)
}

/*
	...
		[a b _]
		  [b _ b]
		  	[_ b a]

		[a a _]
		  [a _ b]
		  	[_ b a]

		[b a _]
		  [a _ b]
		  	[_ b b]
*/
