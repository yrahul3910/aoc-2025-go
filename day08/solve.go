package day08

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"yrahul3910/aoc-2025-go/utils"
)

type Coordinate struct {
	x int64
	y int64
	z int64
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d, %d)", c.x, c.y, c.z)
}

type PointSet []Coordinate
type Clusters []PointSet

func ParseInput(input string) PointSet {
	arr := make(PointSet, 0)

	for line := range strings.SplitSeq(input, "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		parts := strings.Split(line, ",")

		x, _ := strconv.ParseInt(parts[0], 10, 64)
		y, _ := strconv.ParseInt(parts[1], 10, 64)
		z, _ := strconv.ParseInt(parts[2], 10, 64)

		arr = append(arr, Coordinate{x, y, z})
	}

	return arr
}

func InitializeClusters(points PointSet) Clusters {
	clusters := make(Clusters, 0)

	for _, p := range points {
		curCluster := make(PointSet, 1)
		curCluster[0] = p

		clusters = append(clusters, curCluster)
	}

	return clusters
}

func EuclideanDistance(x, y Coordinate) float64 {
	return math.Sqrt(
		float64((x.x-y.x)*(x.x-y.x)) +
			float64((x.y-y.y)*(x.y-y.y)) +
			float64((x.z-y.z)*(x.z-y.z)),
	)
}

func GetDistanceMatrix(points PointSet) [][]float64 {
	n := len(points)

	dist := make([][]float64, n)
	for i := range n {
		dist[i] = make([]float64, n)
	}

	for i := range n {
		for j := range n {
			if i < j {
				continue
			}
			if i == j {
				dist[i][j] = 0.0
			}

			dist[i][j] = EuclideanDistance(points[i], points[j])
			dist[j][i] = EuclideanDistance(points[i], points[j])
		}
	}

	return dist
}

type Index2D struct {
	i int
	j int
}

func GetNLowest(arr [][]float64, n int) []Index2D {
	size := len(arr)

	type Element struct {
		val float64
		idx Index2D
	}

	values := make([]Element, 0)
	for i := range size {
		for j := range size {
			if i < j {
				values = append(values, Element{arr[i][j], Index2D{i, j}})
			}
		}
	}

	slices.SortFunc(values, func(a, b Element) int {
		if a.val != b.val {
			return int((a.val - b.val) / math.Abs(a.val-b.val))
		}

		if a.idx.i != b.idx.i {
			return a.idx.i - b.idx.i
		}

		return a.idx.j - b.idx.j
	})

	result := make([]Index2D, min(n, len(values)))
	for i := range min(n, len(values)) {
		result[i] = values[i].idx
	}

	return result
}

func SolvePuzzle1(input string) int {
	points := ParseInput(input)

	uf := NewUnionFind[Coordinate]()
	for _, point := range points {
		uf.MakeSet(point)
	}

	dist := GetDistanceMatrix(points)
	lowest := GetNLowest(dist, 1000)
	utils.PrintArray(lowest)

	for _, idx := range lowest {
		// merge nodes `idx.i` and `idx.j`
		uf.UnionValues(points[idx.i], points[idx.j])
	}

	sizes := make([]int, 0)
	components := uf.Components()

	for _, component := range components {
		sizes = append(sizes, component.Count)
		uf.PrintComponent(*component)
		fmt.Println()
	}

	slices.Sort(sizes)
	slices.Reverse(sizes)

	return sizes[0] * sizes[1] * sizes[2]
}

func SolvePuzzle2(input string) int {
	points := ParseInput(input)

	uf := NewUnionFind[Coordinate]()
	for _, point := range points {
		uf.MakeSet(point)
	}

	dist := GetDistanceMatrix(points)
	lowest := GetNLowest(dist, 10000) // abuse the `min` in `GetNLowest`

	lastMerged1, lastMerged2 := int64(-1), int64(-1)

	for i := 0; i < len(lowest) && len(uf.Components()) > 1; i++ {
		idx := lowest[i]

		// merge nodes `idx.i` and `idx.j`
		uf.UnionValues(points[idx.i], points[idx.j])

		// only need product of x-coordinates
		lastMerged1 = points[idx.i].x
		lastMerged2 = points[idx.j].x
	}

	return int(lastMerged1) * int(lastMerged2)
}
