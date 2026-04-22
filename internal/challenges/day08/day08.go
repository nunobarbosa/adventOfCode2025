package day08

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Coordinates struct {
	X int
	Y int
	Z int
}

type Box struct {
	ID       int
	Position Coordinates
}

type BoxesDistance struct {
	Box1     Box
	Box2     Box
	Distance int64
}

func Solve(part int, input []string) (string, error) {
	numOfCoordinates, distances, err := BuildInput(input)
	if err != nil {
		return "", err
	}

	switch part {
	case 1:
		return solvePart1(numOfCoordinates, distances)
	case 2:
		return solvePart2(numOfCoordinates, distances)
	default:
		return "", nil
	}
}

func BuildInput(input []string) (int, []BoxesDistance, error) {
	boxes := []Box{}

	for id, l := range input {
		rawValues := strings.Split(l, ",")
		if len(rawValues) != 3 {
			return 0, nil, fmt.Errorf("invalid coordinate %q", l)
		}

		x, err := strconv.Atoi(rawValues[0])
		if err != nil {
			return 0, nil, err
		}

		y, err := strconv.Atoi(rawValues[1])
		if err != nil {
			return 0, nil, err
		}

		z, err := strconv.Atoi(rawValues[2])
		if err != nil {
			return 0, nil, err
		}

		boxes = append(boxes, Box{
			ID:       id,
			Position: Coordinates{X: x, Y: y, Z: z},
		})
	}

	distances := []BoxesDistance{}
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			box1 := boxes[i]
			box2 := boxes[j]
			if box1.compare(box2) > 0 {
				box1, box2 = box2, box1
			}

			d := boxes[i].Position.distance(boxes[j].Position)
			distances = append(distances, BoxesDistance{
				Box1:     box1,
				Box2:     box2,
				Distance: d,
			})
		}
	}

	slices.SortFunc(distances, func(d1, d2 BoxesDistance) int {
		if result := cmp.Compare(d1.Distance, d2.Distance); result != 0 {
			return result
		}

		if result := d1.Box1.compare(d2.Box1); result != 0 {
			return result
		}

		return d1.Box2.compare(d2.Box2)
	})

	return len(boxes), distances, nil
}

func (b Box) compare(b1 Box) int {
	if result := b.Position.compare(b1.Position); result != 0 {
		return result
	}

	return cmp.Compare(b.ID, b1.ID)
}

func (c Coordinates) compare(c1 Coordinates) int {
	if result := cmp.Compare(c.X, c1.X); result != 0 {
		return result
	}

	if result := cmp.Compare(c.Y, c1.Y); result != 0 {
		return result
	}

	return cmp.Compare(c.Z, c1.Z)
}

func (c Coordinates) distance(c1 Coordinates) int64 {
	deltaX := int64(c1.X - c.X)
	deltaY := int64(c1.Y - c.Y)
	deltaZ := int64(c1.Z - c.Z)

	return deltaX*deltaX + deltaY*deltaY + deltaZ*deltaZ
}

// Part 1

// Uses Kruskal's MST algorithm but focusing only on the N shortests pairs.
// For the tests, it uses 10; for the actual input, it uses 1000

const (
	part1ConnectionCount        = 1000
	examplePart1ConnectionCount = 10
	examplePart1Boxes           = 20
)

func solvePart1(numOfCoordinates int, distances []BoxesDistance) (string, error) {
	connectionCount := part1ConnectionTarget(numOfCoordinates, len(distances))
	junctions := newDisjointSet(numOfCoordinates)

	for _, d := range distances[:connectionCount] {
		junctions.union(d.Box1.ID, d.Box2.ID)
	}

	counters := componentSizes(junctions)

	return strconv.Itoa(counters[0] * counters[1] * counters[2]), nil
}

func part1ConnectionTarget(numOfCoordinates, numOfDistances int) int {
	target := part1ConnectionCount
	if numOfCoordinates == examplePart1Boxes {
		target = examplePart1ConnectionCount
	}

	if target > numOfDistances {
		return numOfDistances
	}

	return target
}

// Part 2

// Typical Kruskal's algorithm to build a MST and get the last junction box that creates
// a cycle

func solvePart2(numOfCoordinates int, distances []BoxesDistance) (string, error) {
	if numOfCoordinates <= 1 {
		return "0", nil
	}

	junctions := newDisjointSet(numOfCoordinates)
	for _, d := range distances {
		if !junctions.union(d.Box1.ID, d.Box2.ID) {
			continue
		}

		if junctions.components == 1 {
			result := int64(d.Box1.Position.X) * int64(d.Box2.Position.X)
			return strconv.FormatInt(result, 10), nil
		}
	}

	return "", errors.New("could not connect all junction boxes into one circuit")
}

// Shared Helpers

func componentSizes(junctions *disjointSet) []int {
	counters := map[int]int{}
	for i := range junctions.parent {
		root := junctions.find(i)
		counters[root]++
	}

	sizes := make([]int, 0, len(counters))
	for _, count := range counters {
		sizes = append(sizes, count)
	}

	slices.SortFunc(sizes, func(size1, size2 int) int {
		return cmp.Compare(size2, size1)
	})

	return sizes
}

type disjointSet struct {
	parent     []int
	size       []int
	components int
}

func newDisjointSet(size int) *disjointSet {
	parent := make([]int, size)
	sizes := make([]int, size)

	for i := range size {
		parent[i] = i
		sizes[i] = 1
	}

	return &disjointSet{
		parent:     parent,
		size:       sizes,
		components: size,
	}
}

func (d *disjointSet) find(box int) int {
	if d.parent[box] != box {
		d.parent[box] = d.find(d.parent[box])
	}

	return d.parent[box]
}

func (d *disjointSet) union(box1, box2 int) bool {
	root1 := d.find(box1)
	root2 := d.find(box2)

	if root1 == root2 {
		return false
	}

	if d.size[root1] < d.size[root2] {
		root1, root2 = root2, root1
	}

	d.parent[root2] = root1
	d.size[root1] += d.size[root2]
	d.components--

	return true
}
