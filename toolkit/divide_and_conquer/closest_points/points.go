package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

var eps = 0.001

func (p1 Point) Distance(p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.x)-float64(p2.x), 2) + math.Pow(float64(p1.y)-float64(p2.y), 2))
}

func ReadInput() (int, []Point) {
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Wrong input in 1 line")
		os.Exit(1)
	}

	splitedInput := strings.Split(strings.TrimSpace(str), " ")
	var n int

	n, err = strconv.Atoi(splitedInput[0])
	if err != nil {
		fmt.Println("n param is incorrect")
		os.Exit(1)
	}

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Wrong input for %d spice\n", i)
			os.Exit(1)
		}
		splitedInput := strings.Split(strings.TrimSpace(str), " ")
		points[i].x, err = strconv.Atoi(splitedInput[0])
		if err != nil {
			fmt.Printf("value param is incorrect for spice %d", i)
			os.Exit(1)
		}
		points[i].y, err = strconv.Atoi(splitedInput[1])
		if err != nil {
			fmt.Printf("weight param is incorrect for spice %d", i)
			os.Exit(1)
		}
	}

	return n, points
}

func QuickSort(points []Point, compare func(Point, Point) bool) {
	length := len(points)
	if length <= 1 {
		return
	}
	// fmt.Println(length)
	pv := rand.Intn(length)
	points[pv], points[length-1] = points[length-1], points[pv]
	// fmt.Printf("Sorting %v\n", numbers)
	i := -1
	twin := -1
	for j := 0; j <= length-1; j++ {
		if compare(points[j], points[length-1]) {
			i++
			if j > i {
				points[i], points[j] = points[j], points[i]
			}
			if points[i] != points[length-1] {
				twin++
				if i > twin {
					points[i], points[twin] = points[twin], points[i]
				}
			}
		}
	}
	twin++
	// fmt.Printf("Result %v\n", numbers)

	if twin > 0 {
		QuickSort(points[0:twin], compare)
	}
	if i < length-1 {
		QuickSort(points[i+1:], compare)
	}
}

func binSearch(points []Point, val float64) int {
	l, r, mid := 0, len(points)-1, 0

	for l < r {
		mid = (l + r + 1) / 2
		if math.Abs(val-float64(points[mid].x)) <= eps {
			return mid
		} else if val < float64(points[mid].x) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

func filterPoints(points []Point, d float64, midX int) []Point {
	if d == 0 {
		res := make([]Point, len(points))
		copy(res, points)
		return points
	}
	leftEdge, rightEdge := float64(midX)-d, float64(midX)+d
	l, r := binSearch(points, leftEdge), binSearch(points, rightEdge)
	// fmt.Println(points)
	if l < len(points) {
		if float64(points[l].x) >= leftEdge {
			for l > 0 {
				if points[l-1] != points[l] {
					break
				}
				l--
			}
		} else {
			for l < len(points)-1 {
				if points[l+1] != points[l] {
					break
				}
				l++
			}
			l++
			if l == len(points) {
				return make([]Point, 0)
			}
		}
	}

	if r < len(points) {
		if float64(points[r].x) > rightEdge {
			for r > 0 {
				if points[r-1] != points[r] {
					break
				}
				r--
			}
			r--
			if r < 0 {
				return make([]Point, 0)
			}
		} else {
			for r < len(points)-1 {
				if points[r+1] != points[r] {
					break
				}
				r++
			}

		}
	}

	res := make([]Point, r-l+1)
	copy(res, points[l:r+1])
	return res
}

func getMinDist(d1, d2 float64) float64 {
	if d1 < 0 {
		return d2
	}
	if d2 < 0 {
		return d1
	}
	return math.Min(d1, d2)
}

func CalculateMinDistance(n int, points []Point) float64 {
	if n == 1 {
		return -1
	}
	if n == 2 {
		return points[0].Distance(points[1])
	}

	d1 := CalculateMinDistance(n/2, points[:n/2])
	d2 := CalculateMinDistance(n-n/2, points[n/2:])

	d := getMinDist(d1, d2)
	// fmt.Printf("d = %f\n", d)

	dRange := filterPoints(points, d, 1-n/2)
	if len(dRange) < 2 {
		return d
	}

	QuickSort(dRange, func(p1, p2 Point) bool {
		return p1.y <= p2.y
	})

	dCross := dRange[0].Distance(dRange[1])
	for i, v := range dRange {
		for j := i + 1; j <= i+7 && j < len(dRange); j++ {
			distance := v.Distance(dRange[j])
			if dCross > distance {
				dCross = distance
			}
		}
	}
	// fmt.Printf("dCross = %f\n", dCross)

	return getMinDist(d, dCross)

}

func main() {
	n, points := ReadInput()
	QuickSort(points, func(p1, p2 Point) bool {
		return p1.x <= p2.x
	})
	fmt.Printf("%.6f", CalculateMinDistance(n, points))
	// fmt.Println(strings.Trim(fmt.Sprint(coord), "[]"))
}
