package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Pair struct {
	open_pos, close_pos int
}

type Pairs []Pair

func duplicateKey(data string) (map[string]int, error) {
	st := strings.Fields(data)
	wc := make(map[string]int)
	for _, word := range st {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
			return wc, fmt.Errorf("duplicate key %s", word)
		} else {
			wc[word] = 1
		}
	}

	return wc, nil
}

func find_last(sub_str string) (key string, err error) {
	m := reverse(sub_str)

	pos := strings.IndexRune(m, '"')
	if pos == -1 {
		return "", fmt.Errorf("Not a valid JSON key")
	}

	// key
	key = reverse(m[:pos])
	return
}

func reverse(str string) (s string) {
	for _, r := range str {
		s = string(r) + s
	}
	return
}

func inBounds(current_pos int, pairs Pairs) (int, error) {

	if len(pairs) > 0 {
		// check if key in sub-obj
		pp := pairs[1:]
		for i, p := range pp {
			o, c := p.open_pos, p.close_pos
			if (current_pos >= o) && (current_pos <= c) {
				return i + 1, nil
			}
		}
	}

	// key in root of object
	if (current_pos >= pairs[0].open_pos) && (current_pos <= pairs[0].close_pos) {
		return 0, nil
	}

	return -1, fmt.Errorf("out of bounds")
}

func DetectDuplicateKeys(data []byte) (keys map[string]int, err error) {
	keys = make(map[string]int)
	decoder := json.NewDecoder(bytes.NewReader(data))
	_, err = decoder.Token() // check for {}[]

	if err != nil {
		return
	}

	// strategy - load all open brackets into open array and all closing brackets into close array
	// reverse the close array to align the brackets
	// if the lengths don't match then return an error

	st_d := string(data)
	open_delims := make([]int, 0)
	close_delims := make([]int, 0)

	potential_keys := make([]int, 0)
	pairs := make([]Pair, 0)

	for i := 0; i < len(st_d); i++ {
		c := st_d[i]
		if c == '{' {
			open_delims = append(open_delims, i)
		} else if c == '}' {
			close_delims = append(close_delims, i)
		} else if c == ':' {
			potential_keys = append(potential_keys, i)
			// fmt.Printf("Found potential key at position %d\n", i)
		}
	}

	if len(open_delims) != len(close_delims) {
		err = fmt.Errorf("invalid JSON object - missing open/closing bracket")
		return
	}

	// first open last closed
	// open[second:] -> closed[:last]

	// reverse the open delim array to align with the closing delim array
	// we want to start with the smallest embedded object first then percolate upwards
	// https://github.com/golang/go/wiki/SliceTricks#reversing
	//for i := len(close_delims)/2 - 1; i >= 0; i-- {
	//	opp := len(close_delims) - 1 - i
	//	close_delims[i], close_delims[opp] = close_delims[opp], close_delims[i]
	//}

	first_open := open_delims[0]
	open_delims = open_delims[1:]

	last_close := close_delims[len(close_delims)-1]
	close_delims = close_delims[:len(close_delims)-1]

	pairs = append(pairs, Pair{first_open, last_close})

	for i := 0; i < len(close_delims); i++ {
		pairs = append(pairs, Pair{open_delims[i], close_delims[i]})
	}

	// for each potential key we want to confirm that we have "key_name" prior to the ':'
	for i := 0; i < len(potential_keys)-1; i++ {
		// check ':' wasn't at index zero
		if potential_keys[i]-1 <= 0 {
			// return error
			err = fmt.Errorf("Invalid JSON object - missing object key prior to ':'")
			return
		} else if st_d[potential_keys[i]-1] != '"' {
			// not a key
			// if url ignore else return error
			// return fmt.Errorf("Invalid JSON object - missing object key prior to ':'")
		} else {
			// found '"' search for the one prior to it
			// omitting ": from the end of the sub string to search
			pkey := potential_keys[i]
			sub_str := st_d[:pkey-1]
			key, _ := find_last(sub_str)
			// which object is this key in?
			level, err := inBounds(pkey, pairs)
			if err != nil {
				return keys, err
			}
			// fmt.Printf("%s %d\n", key, level)
			if _, ok := keys[fmt.Sprintf("%s_%d", key, level)]; ok {
				keys[fmt.Sprintf("%s_%d", key, level)] += 1
			} else {
				keys[fmt.Sprintf("%s_%d", key, level)] = 1
			}
		}
	}
	return
}
