package day20

import (
	"regexp"
	"strconv"

	"fmt"

	"github.com/AlbertBrand/aoc2017/util"
	"golang.org/x/tools/container/intsets"
)

type Particle struct {
	x, y, z    int
	vx, vy, vz int
	ax, ay, az int
	collided   bool
}

func (p *Particle) tick() {
	p.vx += p.ax
	p.vy += p.ay
	p.vz += p.az

	p.x += p.vx
	p.y += p.vy
	p.z += p.vz
}

func (p *Particle) dist() int {
	return abs(p.x) + abs(p.y) + abs(p.z)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

var re = regexp.MustCompile(`p=<(-?\d+),(-?\d+),(-?\d+)>, v=<(-?\d+),(-?\d+),(-?\d+)>, a=<(-?\d+),(-?\d+),(-?\d+)>`)

func parseLine(line string) Particle {
	result := re.FindStringSubmatch(line)
	x, _ := strconv.Atoi(result[1])
	y, _ := strconv.Atoi(result[2])
	z, _ := strconv.Atoi(result[3])

	vx, _ := strconv.Atoi(result[4])
	vy, _ := strconv.Atoi(result[5])
	vz, _ := strconv.Atoi(result[6])

	ax, _ := strconv.Atoi(result[7])
	ay, _ := strconv.Atoi(result[8])
	az, _ := strconv.Atoi(result[9])

	return Particle{
		x, y, z, vx, vy, vz, ax, ay, az, false,
	}
}

func solver(lines []string) {
	particles := make([]*Particle, 0)
	for _, line := range lines {
		particle := parseLine(line)
		particles = append(particles, &particle)
	}

	for i := 0; i < 1000; i++ {
		for _, particle := range particles {
			particle.tick()
		}
	}

	id := -1
	min := intsets.MaxInt
	for k, particle := range particles {
		dist := particle.dist()
		if dist < min {
			id = k
			min = dist
		}
	}
	println(id)
}

func solver2(lines []string) {
	particles := make([]*Particle, 0)
	for _, line := range lines {
		particle := parseLine(line)
		particles = append(particles, &particle)
	}

	for i := 0; i < 1000; i++ {
		for _, particle := range particles {
			particle.tick()
		}
		collisionMap := make(map[string]*Particle)
		for _, p := range particles {
			if p.collided {
				continue
			}
			key := fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
			if c, exists := collisionMap[key]; !exists {
				collisionMap[key] = p
			} else {
				c.collided = true
				p.collided = true
			}
		}
	}

	count := 0
	for _, particle := range particles {
		if !particle.collided {
			count++
		}
	}
	println(count)
}

func TestFirst() {
	solver(util.ReadTxt("day20/test.txt"))
}

func SolveFirst() {
	solver(util.ReadTxt("day20/input.txt"))
}

func TestSecond() {
	solver2(util.ReadTxt("day20/test2.txt"))
}

func SolveSecond() {
	solver2(util.ReadTxt("day20/input.txt"))
}
