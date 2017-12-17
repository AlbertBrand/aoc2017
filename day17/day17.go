package day17

func solver(steps int) {
	state := make([]int, 1, 2018)
	state[0] = 0
	insertpos := 0
	for i := 1; i < 2018; i++ {
		insertpos = ((insertpos + steps) % i) + 1
		state = append(state, 0)
		copy(state[insertpos+1:], state[insertpos:])
		state[insertpos] = i
	}
	println(state[insertpos+1])
}

func solver2(steps uint32) {
	var insertpos uint32
	var after uint32
	for i := uint32(1); i < 50000000; i++ {
		insertpos = ((insertpos + steps) % i) + 1
		if insertpos == 1 {
			after = i
		}
	}
	println(after)
}

func solver3(steps uint32) {
	var insertpos uint32
	var after uint32
	var i uint32
	for {
		fits := (i - insertpos) / steps
		i += fits + 1
		insertpos = (insertpos+(fits+1)*(steps+1)-1)%i + 1
		if insertpos == 1 {
			after = i
		}
		if i >= 50000000 {
			break
		}
	}
	println(after)
}

func TestFirst() {
	solver(3)
}

func SolveFirst() {
	solver(335)
}

func SolveSecond() {
	solver3(335)
}
