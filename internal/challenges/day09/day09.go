package day09

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Tile struct {
	X int
	Y int
}

type Rectangle struct {
	TopLeft     Tile
	BottomRight Tile
}

type AllowedRectangles struct {
	xIndex        map[int]int
	yIndex        map[int]int
	invalidPrefix [][]int
}

func Solve(part int, input []string) (string, error) {
	tiles, err := BuildInput(input)
	if err != nil {
		return "", err
	}
	switch part {
	case 1:
		return solvePart1(tiles)
	case 2:
		return solvePart2(tiles)
	default:
		return "", nil
	}
}

func BuildInput(input []string) ([]Tile, error) {
	tiles := []Tile{}
	for _, l := range input {
		rawValues := strings.Split(l, ",")
		if len(rawValues) != 2 {
			return nil, fmt.Errorf("invalid coordinate %q", l)
		}

		x, err := strconv.Atoi(rawValues[0])
		if err != nil {
			return nil, err
		}

		y, err := strconv.Atoi(rawValues[1])
		if err != nil {
			return nil, err
		}

		tiles = append(tiles, Tile{X: x, Y: y})
	}

	return tiles, nil
}

// Part 1

func solvePart1(tiles []Tile) (string, error) {
	maxArea := 0

	for i := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			area := tiles[i].area(tiles[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return strconv.Itoa(maxArea), nil
}

func solvePart2(redTiles []Tile) (string, error) {
	if len(redTiles) == 0 {
		return "0", nil
	}

	allowed, err := buildAllowedRectangles(redTiles)
	if err != nil {
		return "", err
	}

	area := 0
	for i := range redTiles {
		for j := i + 1; j < len(redTiles); j++ {
			rectangle := buildReactangle(redTiles[i], redTiles[j])
			if allowed.contains(rectangle) {
				area = max(area, rectangle.area())
			}
		}
	}

	return strconv.Itoa(area), nil
}

// Helpers

func (t Tile) area(t1 Tile) int {
	minX := t.X
	maxX := t1.X
	if minX > maxX {
		minX, maxX = maxX, minX
	}

	minY := t.Y
	maxY := t1.Y
	if minY > maxY {
		minY, maxY = maxY, minY
	}

	return (maxX - minX + 1) * (maxY - minY + 1)
}

func buildReactangle(t1, t2 Tile) Rectangle {
	return Rectangle{
		TopLeft: Tile{
			X: min(t1.X, t2.X),
			Y: min(t1.Y, t2.Y),
		},
		BottomRight: Tile{
			X: max(t1.X, t2.X),
			Y: max(t1.Y, t2.Y),
		},
	}
}

func (r Rectangle) area() int {
	return r.TopLeft.area(r.BottomRight)
}

func buildAllowedRectangles(redTiles []Tile) (AllowedRectangles, error) {
	xValues := uniqueSortedCoordinates(redTiles, func(tile Tile) int { return tile.X })
	yValues := uniqueSortedCoordinates(redTiles, func(tile Tile) int { return tile.Y })

	xIndex, width := compressAxis(xValues)
	yIndex, height := compressAxis(yValues)

	boundary := make([][]bool, height)
	for y := range boundary {
		boundary[y] = make([]bool, width)
	}

	for i, current := range redTiles {
		next := redTiles[(i+1)%len(redTiles)]

		switch {
		case current.X == next.X:
			x := xIndex[current.X]
			startY := min(yIndex[current.Y], yIndex[next.Y])
			endY := max(yIndex[current.Y], yIndex[next.Y])
			for y := startY; y <= endY; y++ {
				boundary[y][x] = true
			}
		case current.Y == next.Y:
			y := yIndex[current.Y]
			startX := min(xIndex[current.X], xIndex[next.X])
			endX := max(xIndex[current.X], xIndex[next.X])
			for x := startX; x <= endX; x++ {
				boundary[y][x] = true
			}
		default:
			return AllowedRectangles{}, fmt.Errorf("invalid edge from %v to %v", current, next)
		}
	}

	outside := floodFillOutside(boundary)
	invalidPrefix := makePrefixSum(boundary, outside)

	return AllowedRectangles{
		xIndex:        xIndex,
		yIndex:        yIndex,
		invalidPrefix: invalidPrefix,
	}, nil
}

func uniqueSortedCoordinates(tiles []Tile, selector func(Tile) int) []int {
	valuesMap := map[int]bool{}
	for _, tile := range tiles {
		valuesMap[selector(tile)] = true
	}

	values := make([]int, 0, len(valuesMap))
	for value := range valuesMap {
		values = append(values, value)
	}
	slices.Sort(values)

	return values
}

func compressAxis(values []int) (map[int]int, int) {
	index := map[int]int{}
	segments := 0

	for i, value := range values {
		index[value] = segments
		segments++

		if i == len(values)-1 {
			continue
		}

		if values[i+1]-value > 1 {
			segments++
		}
	}

	return index, segments
}

func floodFillOutside(boundary [][]bool) [][]bool {
	height := len(boundary)
	width := len(boundary[0])

	outside := make([][]bool, height)
	for y := range outside {
		outside[y] = make([]bool, width)
	}

	queue := make([]Tile, 0, width*2+height*2)
	push := func(x, y int) {
		if boundary[y][x] || outside[y][x] {
			return
		}

		outside[y][x] = true
		queue = append(queue, Tile{X: x, Y: y})
	}

	for x := range width {
		push(x, 0)
		push(x, height-1)
	}
	for y := 1; y < height-1; y++ {
		push(0, y)
		push(width-1, y)
	}

	for head := 0; head < len(queue); head++ {
		current := queue[head]
		if current.X > 0 {
			push(current.X-1, current.Y)
		}
		if current.X+1 < width {
			push(current.X+1, current.Y)
		}
		if current.Y > 0 {
			push(current.X, current.Y-1)
		}
		if current.Y+1 < height {
			push(current.X, current.Y+1)
		}
	}

	return outside
}

func makePrefixSum(boundary, outside [][]bool) [][]int {
	height := len(boundary)
	width := len(boundary[0])

	prefix := make([][]int, height+1)
	for y := range prefix {
		prefix[y] = make([]int, width+1)
	}

	for y := range height {
		for x := range width {
			invalid := 0
			if !boundary[y][x] && outside[y][x] {
				invalid = 1
			}

			prefix[y+1][x+1] = invalid + prefix[y][x+1] + prefix[y+1][x] - prefix[y][x]
		}
	}

	return prefix
}

func (a AllowedRectangles) contains(rectangle Rectangle) bool {
	startX := a.xIndex[rectangle.TopLeft.X]
	endX := a.xIndex[rectangle.BottomRight.X]
	startY := a.yIndex[rectangle.TopLeft.Y]
	endY := a.yIndex[rectangle.BottomRight.Y]

	invalid := a.invalidPrefix[endY+1][endX+1] -
		a.invalidPrefix[startY][endX+1] -
		a.invalidPrefix[endY+1][startX] +
		a.invalidPrefix[startY][startX]

	return invalid == 0
}
