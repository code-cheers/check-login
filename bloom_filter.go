package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/bits-and-blooms/bloom/v3"
)

func main() {
	inPath := flag.String("in", "data/user_ids.txt", "input file path")
	falsePosPct := flag.Float64("false-pos", 0.001, "false positive rate")
	queries := flag.Int("queries", 100_000, "number of hit/miss queries")
	flag.Parse()

	ids, err := readIDs(*inPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read IDs failed: %v\n", err)
		os.Exit(1)
	}
	if len(ids) == 0 {
		fmt.Fprintln(os.Stderr, "no IDs loaded")
		os.Exit(1)
	}

	bf := bloom.NewWithEstimates(uint(len(ids)), *falsePosPct)

	startInsert := time.Now()
	for _, id := range ids {
		bf.AddString(id)
	}
	insertDur := time.Since(startInsert)

	hitQueries := *queries
	if hitQueries > len(ids) {
		hitQueries = len(ids)
	}

	startHit := time.Now()
	for i := 0; i < hitQueries; i++ {
		if !bf.TestString(ids[i]) {
			panic("expected hit, got miss")
		}
	}
	hitDur := time.Since(startHit)

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	startMiss := time.Now()
	falsePos := 0
	for i := 0; i < *queries; i++ {
		key := usernameForID(len(ids) + rng.Intn(len(ids)))
		if bf.TestString(key) {
			falsePos++
		}
	}
	missDur := time.Since(startMiss)

	fmt.Printf("items=%d, fp_est=%.2f%%\n", len(ids), *falsePosPct*100)
	fmt.Printf("insert=%s, hit_check=%s, miss_check=%s\n", insertDur, hitDur, missDur)
	fmt.Printf("false_pos=%d/%d (%.2f%%)\n", falsePos, *queries, float64(falsePos)/float64(*queries)*100)
}

func readIDs(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 64), 1_000_000)
	ids := make([]string, 0, 1_000_000)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return ids, nil
}

var givenNames = []string{
	"ming", "wei", "jun", "hao", "lei", "yan", "ting", "qi", "yu", "lin", "jie", "xin",
}

var familyNames = []string{
	"zhang", "wang", "li", "liu", "chen", "yang", "huang", "zhou", "wu", "xu", "sun", "zhao",
}

func usernameForID(id int) string {
	given := givenNames[id%len(givenNames)]
	family := familyNames[(id/len(givenNames))%len(familyNames)]
	return fmt.Sprintf("%s.%s%06d", given, family, id)
}
