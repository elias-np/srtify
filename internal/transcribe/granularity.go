package transcribe

import "strconv"

// whisper.cpp has no native "words per segment" option. Its -ml flag caps
// segment length in characters, and -sow (split-on-word) rounds that cut to
// the nearest word boundary instead of the nearest token. We approximate a
// words-per-segment target using ~6 characters per word (5 letters + 1
// space, a reasonable average for pt/en speech). Level 5 uses -ml 1, which
// forces a break after every word since no word fits a 1-character budget
// and -sow always rounds up to a full word.
var maxLenByGranularity = map[int]int{
	1: 60, // long phrases, ~10 words
	2: 42, // medium phrases, ~7 words
	3: 30, // short phrases, ~5 words
	4: 20, // small groups, ~3-5 words
	5: 1,  // word by word
}

func granularityArgs(level int) []string {
	maxLen, ok := maxLenByGranularity[level]
	if !ok {
		return nil
	}

	return []string{"-ml", strconv.Itoa(maxLen), "-sow"}
}
