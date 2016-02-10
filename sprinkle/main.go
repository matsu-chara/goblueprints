package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/kardianos/osext"
)

const otherWord = "*"

type transforms struct {
	Words []string `json:"words"`
}

func (t transforms) pickWord(r *rand.Rand) string {
	return t.Words[r.Intn(len(t.Words))]
}

func readTransforms(filename string) (transforms, error) {
	file, err := os.Open(filename)
	if err != nil {
		return transforms{}, err
	}
	decoder := json.NewDecoder(file)

	var t transforms
	err2 := decoder.Decode(&t)
	if err2 != nil {
		return transforms{}, err
	}
	return t, nil
}

func main() {
	filedir, err := osext.ExecutableFolder()
	if err != nil {
		fmt.Printf("get file dir error: %v", err)
	}

	t, err := readTransforms(filedir + "/transforms.json")
	if err != nil {
		fmt.Printf("read transforms error: %v", err)
	}

	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(strings.Replace(t.pickWord(r), otherWord, s.Text(), -1))
	}
}
