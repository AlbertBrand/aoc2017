package day3

type Vector struct {
	x, y int
}

type Square struct {
	Vector
	value int
}

var up = Vector{0, 1}
var left = Vector{-1, 0}
var down = Vector{0, -1}
var right = Vector{1, 0}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (c *Vector) dist() int {
	return abs(c.x) + abs(c.y)
}

func (c *Vector) add(a Vector) {
	c.x += a.x
	c.y += a.y
}

func solver(square int) {
	cursor := 1
	direction := right
	sideLen := 1
	currSideLen := 0
	coordinate := Vector{0, 0}
	for {
		if cursor == square {
			// found
			println(coordinate.dist())
			break
		}

		// check for turn
		if currSideLen == sideLen {
			switch direction {
			case up:
				direction = left
				sideLen++
			case left:
				direction = down
			case down:
				direction = right
				sideLen++
			case right:
				direction = up
			}
			currSideLen = 0
		}

		// travel
		coordinate.add(direction)
		currSideLen++
		cursor++
	}
}

func solver2(minimum int) {
	counter := 1
	direction := right
	sideLen := 1
	currSideLen := 0
	coordinate := Vector{0, 0}

	squares := make([]Square, 0)

	for {
		// found?
		if counter > minimum {
			println(counter)
			break
		}

		// add to squares list
		squares = append(squares, Square{
			Vector: coordinate,
			value:  counter,
		})

		// check for turn
		if currSideLen == sideLen {
			switch direction {
			case up:
				direction = left
				sideLen++
			case left:
				direction = down
			case down:
				direction = right
				sideLen++
			case right:
				direction = up
			}
			currSideLen = 0
		}

		// travel
		coordinate.add(direction)
		currSideLen++

		// calculate new counter
		counter = 0
		for _, square := range squares {
			if square.x >= coordinate.x-1 &&
				square.x <= coordinate.x+1 &&
				square.y >= coordinate.y-1 &&
				square.y <= coordinate.y+1 {
				counter += square.value
			}
		}
	}
}

func TestFirst() {
	solver(1)
	solver(12)
	solver(23)
	solver(1024)
}

func SolveFirst() {
	solver(289326)
}

func TestSecond() {
	solver2(147)
	solver2(747)
}

func SolveSecond() {
	solver2(289326)
}
