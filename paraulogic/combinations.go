package paraulogic

import "math/rand/v2"

func randIntInRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func selectRandomlyFromRunes(pool []rune, n int) []rune {
	if n > len(pool) {
		panic("not enough unique elements in pool to select")
	}

	// Make a copy to avoid mutating the original slice
	shuffled := make([]rune, len(pool))
	copy(shuffled, pool)

	// Shuffle using Fisher-Yates (in-place)
	for i := len(shuffled) - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled[:n]
}

func permute(runes []rune, minLength int, maxLength int) [][]rune {
	var result [][]rune

	var backtrack func(path []rune)
	backtrack = func(path []rune) {
		if len(path) > minLength {
			perm := make([]rune, len(path))
			copy(perm, path)
			result = append(result, perm)
		}
		if len(path) == maxLength {
			return
		}
		for _, r := range runes {
			path = append(path, r)
			backtrack(path)
			path = path[:len(path)-1]
		}
	}

	backtrack([]rune{})
	return result
}
