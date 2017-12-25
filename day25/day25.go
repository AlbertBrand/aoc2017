package day25

type State uint
const (
	A State = iota
	B
	C
	D
	E
	F
)

func TestFirst() {
	state := A
	pos := 0
	slot := make(map[int]bool)
	for i:=0; i<6; i++ {
		switch state {
		case A:
			if slot[pos] == false {
				slot[pos] = true
				pos++
				state = B
			} else {
				slot[pos] = false
				pos--
				state = B
			}
		case B:
			if slot[pos] == false {
				slot[pos] = true
				pos--
				state = A
			} else {
				slot[pos] = true
				pos++
				state = A
			}
		}
	}
	check := 0
	for _, val := range slot {
		if val {
			check++
		}
	}
	println(check)
}

func SolveFirst() {
	state := A
	pos := 0
	slot := make(map[int]bool)
	for i:=0; i<12861455; i++ {
		switch state {
		case A:
			if slot[pos] == false {
				slot[pos] = true
				pos++
				state = B
			} else {
				slot[pos] = false
				pos--
				state = B
			}
		case B:
			if slot[pos] == false {
				slot[pos] = true
				pos--
				state = C
			} else {
				slot[pos] = false
				pos++
				state = E
			}
		case C:
			if slot[pos] == false {
				slot[pos] = true
				pos++
				state = E
			} else {
				slot[pos] = false
				pos--
				state = D
			}
		case D:
			if slot[pos] == false {
				slot[pos] = true
				pos--
				state = A
			} else {
				pos--
				state = A
			}
		case E:
			if slot[pos] == false {
				pos++
				state = A
			} else {
				slot[pos] = false
				pos++
				state = F
			}
		case F:
			if slot[pos] == false {
				slot[pos] = true
				pos++
				state = E
			} else {
				pos++
				state = A
			}
		}
	}
	check := 0
	for _, val := range slot {
		if val {
			check++
		}
	}
	println(check)
}