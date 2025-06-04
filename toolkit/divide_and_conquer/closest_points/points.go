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
	if n < 2 {
		return -1
	}
	pointsX := make([]Point, n)
	pointsY := make([]Point, n)
	copy(pointsX, points)
	copy(pointsY, points)
	QuickSort(pointsX, func(p1, p2 Point) bool { return p1.x < p2.x })
	QuickSort(pointsY, func(p1, p2 Point) bool { return p1.y < p2.y })
	return closestUtil(pointsX, pointsY)
}

func closestUtil(pointsX, pointsY []Point) float64 {
	n := len(pointsX)
	if n <= 3 {
		minDist := -1.0
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				d := pointsX[i].Distance(pointsX[j])
				if minDist < 0 || d < minDist {
					minDist = d
				}
			}
		}
		return minDist
	}
	mid := n / 2
	midX := pointsX[mid].x

	leftX := pointsX[:mid]
	rightX := pointsX[mid:]

	leftY := make([]Point, 0, mid)
	rightY := make([]Point, 0, n-mid)
	for _, p := range pointsY {
		if p.x < midX || (p.x == midX && len(leftY) < len(leftX)) {
			leftY = append(leftY, p)
		} else {
			rightY = append(rightY, p)
		}
	}

	d1 := closestUtil(leftX, leftY)
	d2 := closestUtil(rightX, rightY)
	d := getMinDist(d1, d2)

	strip := make([]Point, 0)
	for _, p := range pointsY {
		if math.Abs(float64(p.x-midX)) < d {
			strip = append(strip, p)
		}
	}
	minStrip := d
	for i := 0; i < len(strip); i++ {
		for j := i + 1; j < len(strip) && float64(strip[j].y-strip[i].y) < d; j++ {
			dist := strip[i].Distance(strip[j])
			if dist < minStrip {
				minStrip = dist
			}
		}
	}
	return minStrip
}

func main() {
	n, points := ReadInput()
	QuickSort(points, func(p1, p2 Point) bool {
		return p1.x <= p2.x
	})
	fmt.Printf("%.6f", CalculateMinDistance(n, points))
	// fmt.Println(strings.Trim(fmt.Sprint(coord), "[]"))
}
