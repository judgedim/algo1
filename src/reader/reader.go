package reader

import (
	"log"
	"os"
	"bufio"
	"strconv"
)

func ReadArray(filename string) []int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	array := []int64{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		element, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		array = append(array, element)
	}
	return array
}
