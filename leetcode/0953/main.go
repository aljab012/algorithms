package main

func isAlienSorted(words []string, order string) bool {
	pos := map[byte]int{}
	for i := range order {
		pos[order[i]] = i
	}
	isSorted := func(w1, w2 string) bool {

		for i := 0; i < len(w1) && i < len(w2); i++ {
			if pos[w1[i]] < pos[w2[i]] {
				return true
			}
			if pos[w1[i]] > pos[w2[i]] {
				return false
			}
		}
		return len(w1) <= len(w2)
	}
	for i := 1; i < len(words); i++ {
		if !isSorted(words[i-1], words[i]) {
			return false
		}
	}
	return true
}
