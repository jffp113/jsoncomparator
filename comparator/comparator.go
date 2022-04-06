package comparator

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
)

type parsedFile struct {
	value []map[string]string
	err   error
}

//CompareJSON receives two readers for a array with json objects and
//returns true if both readers contain
func CompareJSON(firstReader io.Reader, secondReader io.Reader) (bool, error) {

	firstCh := parseFile(firstReader)
	secondCh := parseFile(secondReader)

	firstVal := <-firstCh

	if firstVal.err != nil {
		return false, firstVal.err
	}

	secondVal := <-secondCh

	if secondVal.err != nil {
		return false, secondVal.err
	}

	matcher := make(map[string]bool)

	s1 := firstVal.value
	s2 := secondVal.value

	if len(s1) != len(s2) {
		return false, nil
	}

	for _, v := range s1 {
		h := calculateHashFromMap(v)
		matcher[h] = true
	}

	for _, v := range s2 {
		h := calculateHashFromMap(v)
		if !matcher[h] {
			return false, nil
		}
	}

	return true, nil
}

func parseFile(r io.Reader) <-chan parsedFile {

	ch := make(chan parsedFile, 1)

	go func() {
		var j []map[string]string
		bs, err := io.ReadAll(r)

		if err != nil {
			ch <- parsedFile{
				err: err,
			}
		}

		err = json.Unmarshal(bs, &j)

		ch <- parsedFile{
			err:   err,
			value: j,
		}
	}()

	return ch
}

func calculateHashFromMap(toHash map[string]string) string {
	var sb strings.Builder

	keys := make([]string, 0, len(toHash))

	for k := range toHash {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		sb.WriteString(k)
		sb.WriteString(toHash[k])
	}

	return hash(sb.String())
}

func hash(toHash string) string {
	h := sha1.Sum([]byte(toHash))
	return fmt.Sprintf("%x\n", h)
}
