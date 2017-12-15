package day15

type Generator struct {
	factor int
	value  int
	mask   int
}

func (g *Generator) next() {
	g.value *= g.factor
	g.value %= 2147483647
	if g.value&g.mask != 0 {
		g.next()
	}
}

func solver(genA Generator, genB Generator, max int) {
	count := 0
	mask := 1<<16 - 1
	for i := 0; i < max; i++ {
		genA.next()
		genB.next()
		if genA.value&mask == genB.value&mask {
			count++
		}
	}
	println(count)
}

func TestFirst() {
	solver(Generator{16807, 65, 0}, Generator{48271, 8921, 0}, 40000000)
}

func SolveFirst() {
	solver(Generator{16807, 618, 0}, Generator{48271, 814, 0}, 40000000)
}

func TestSecond() {
	solver(Generator{16807, 65, 1<<2 - 1}, Generator{48271, 8921, 1<<3 - 1}, 5000000)
}

func SolveSecond() {
	solver(Generator{16807, 618, 1<<2 - 1}, Generator{48271, 814, 1<<3 - 1}, 5000000)
}
