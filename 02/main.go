package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	num, err := strconv.Atoi(strings.TrimSuffix(string(b), "\n"))
	if err != nil {
		log.Fatalln(err)
	}

	x, y := 0, 0

	if num != 0 {
		// количество полных квадратов
		k := num / 8
		// остаток шагов на не полном квадрате
		o := num % 8
		// количество шагов в стороне не полного квадрата
		L := (num/8 + 1) * 2
		// остаток шагов на неполном квадрате
		r := o - 1
		// l, количество полных сторон
		p := r
		// остаток на неполной стороне
		rL := r % L
		if p == 0 {
			// на первой стороне
			x, y = (-k - 1), (-k + rL)
		} else if p == 1 {
			// на второй стороне
			x, y = (-k - 1), (-k - rL)
		} else if p == 2 {
			// на третьей стороне
			x, y = (-k + rL), (-k - 1)
		} else {
			// на четвертой стороне
			// тут я затупил
			x, y = (-k - 1 + L), (-k + L - 1 - rL)
		}
		log.Println(x, y)
	}

	file, err := os.Create("./output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(x) + " " + strconv.Itoa(y))
	if err != nil {
		log.Fatal(err)
	}
}
