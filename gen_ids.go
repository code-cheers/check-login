package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	count := flag.Int("count", 1_000_000, "number of user IDs to generate")
	outPath := flag.String("out", "data/user_ids.txt", "output file path")
	flag.Parse()

	if *count <= 0 {
		fmt.Fprintln(os.Stderr, "count must be positive")
		os.Exit(1)
	}

	if err := writeIDs(*outPath, *count); err != nil {
		fmt.Fprintf(os.Stderr, "generate IDs failed: %v\n", err)
		os.Exit(1)
	}
}

func writeIDs(path string, count int) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	w := bufio.NewWriter(f)
	for i := 0; i < count; i++ {
		if _, err := w.WriteString(usernameForID(i) + "\n"); err != nil {
			return err
		}
	}
	return w.Flush()
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
