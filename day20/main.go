package main

import (
	"day20/lines"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	tiles := lines.MustParse("data", "\n\n")
	monster := lines.MustParse("monster", "\n")
	start := time.Now()
	var tls Tiles
	for _, tilestr := range tiles {
		tls = append(tls, parseTile(tilestr))
	}
	tls = tls.arrange()
	fmt.Println(part1(tls))
	fmt.Println(part2(tls, monster))
	stop := time.Now()
	fmt.Println(stop.Sub(start))
}

func part1(tls Tiles) int {
	return tls[0].number * tls[tls.size()-1].number * tls[len(tls)-tls.size()].number * tls[len(tls)-1].number
}

func part2(tls Tiles, monster []string) int {
	image := tls.toImage()
	var mimg Image
	for _, row := range monster {
		var r Row
		for _, c := range row {
			r = append(r, c == '#')
		}
		mimg = append(mimg, r)
	}
	monsterCnt := image.findMonsters(mimg)
	result := image.pixelsActive() - monsterCnt*mimg.pixelsActive()
	return result
}

type Row []bool

func (r Row) equal(r2 Row) bool {
	for ix := range r {
		if r[ix] != r2[ix] {
			return false
		}
	}
	return true
}

func (r Row) reverse() Row {
	var result []bool
	for ix := len(r) - 1; ix >= 0; ix-- {
		result = append(result, r[ix])
	}
	return result
}

type Tile struct {
	image         Image
	number        int
	x, y          int
	foundLocation bool
}

func parseTile(tilestr string) Tile {
	tilerows := strings.Split(tilestr, "\n")
	tileNum, _ := strconv.ParseInt(strings.TrimSuffix(strings.TrimPrefix(tilerows[0], "Tile "), ":"), 0, 0)
	var rows []Row
	for _, r := range tilerows[1:] {
		var row Row
		for _, c := range r {
			row = append(row, c == '#')
		}
		rows = append(rows, row)
	}
	return Tile{number: int(tileNum), image: rows}
}

func (t Tile) match(t2 Tile) (*Tile, bool) {
	if !t.foundLocation {
		return nil, false
	}
	if t2.foundLocation {
		return nil, false
	}
	var trot = t2
	for rotation := 0; rotation < 8; rotation++ {
		if rotation%2 == 1 {
			trot.image = trot.image.flipvert()
		}
		if rotation > 0 && rotation%2 == 0 {
			// flip back and rotate 90 degrees
			trot.image = trot.image.flipvert().rotate90()
		}
		if trot.getRow("left").equal(t.getRow("right")) {
			trot.foundLocation = true
			trot.y = t.y
			trot.x = t.x + 1
			return &trot, true
		}
		if trot.getRow("right").equal(t.getRow("left")) {
			trot.foundLocation = true
			trot.y = t.y
			trot.x = t.x - 1
			return &trot, true
		}
		if trot.getRow("top").equal(t.getRow("bottom")) {
			trot.foundLocation = true
			trot.y = t.y + 1
			trot.x = t.x
			return &trot, true
		}
		if trot.getRow("bottom").equal(t.getRow("top")) {
			trot.foundLocation = true
			trot.y = t.y - 1
			trot.x = t.x
			return &trot, true
		}
	}
	return nil, false
}

func (t Tile) getRow(side string) Row {
	var pos int
	switch side {
	case "top":
		return t.image[0]
	case "bottom":
		return t.image[len(t.image)-1]
	case "left":
		pos = 0
	case "right":
		pos = len(t.image[0]) - 1
	}
	var r Row
	for _, row := range t.image {
		r = append(r, row[pos])
	}
	return r
}

func (t Tile) String() string {
	result := fmt.Sprintf("Tile %d\n", t.number)
	result += fmt.Sprint(t.image)
	return result
}

type Tiles []Tile

func (tls Tiles) allFound() bool {
	for _, t := range tls {
		if !t.foundLocation {
			return false
		}
	}
	return true
}

func (tls Tiles) size() int { return int(math.Sqrt(float64(len(tls)))) }

func (tls Tiles) Len() int      { return len(tls) }
func (tls Tiles) Swap(i, j int) { tls[i], tls[j] = tls[j], tls[i] }
func (tls Tiles) Less(i, j int) bool {
	return tls[i].y*tls.size()+tls[i].x < tls[j].y*tls.size()+tls[j].x
}

func (tls Tiles) arrange() Tiles {
	tls[0].foundLocation = true
	for !tls.allFound() {
		for i := 0; i < len(tls); i++ {
			for j := 0; j < len(tls); j++ {
				rotatedTile, match := tls[i].match(tls[j])
				if match {
					tls[j] = *rotatedTile
				}
			}
		}
	}
	sort.Sort(tls)
	return tls
}

func (tls Tiles) toImage() Image {
	var img Image
	size := tls.size()
	tilePixelMax := len(tls[0].image) - 1
	for y := 0; y < size; y++ {
		for rownum := 1; rownum < tilePixelMax; rownum++ {
			var r Row
			for x := 0; x < size; x++ {
				tilenum := y*size + x
				row := tls[tilenum].image[rownum][1:tilePixelMax]
				r = append(r, row...)
			}
			img = append(img, r)
		}
	}
	return img
}

type Image []Row

func (img Image) String() string {
	result := ""
	for _, r := range img {
		for _, v := range r {
			if v {
				result += "#"
			} else {
				result += "."
			}
		}
		result += "\n"
	}
	return result
}

func (img Image) fliphor() Image {
	var newImg Image
	for _, r := range img {
		newImg = append(newImg, r.reverse())
	}
	return newImg
}

func (img Image) flipvert() Image {
	var newImg Image
	for ix := len(img) - 1; ix >= 0; ix-- {
		newImg = append(newImg, img[ix])
	}
	return newImg
}

func (img Image) rotate90() Image {
	var newImg Image
	for i, row := range img {
		var r Row
		for j := range row {
			r = append(r, img[len(img)-j-1][i])
		}
		newImg = append(newImg, r)
	}
	return newImg
}

func (img Image) pxOffsets() [][]int {
	var offsets [][]int
	for y, row := range img {
		for x, active := range row {
			if active {
				offsets = append(offsets, []int{x, y})
			}
		}
	}
	return offsets
}

func (img Image) pixelsActive() int {
	count := 0
	for _, row := range img {
		for _, v := range row {
			if v {
				count++
			}
		}
	}
	return count
}

func (img Image) findMonsters(monster Image) int {
	offsets := monster.pxOffsets()
	xsize := len(monster[0])
	ysize := len(monster)
	imgrot := img
	monsterCount := 0
	for rotation := 0; rotation < 8; rotation++ {
		if rotation%2 == 1 {
			imgrot = imgrot.flipvert()
		}
		if rotation > 0 && rotation%2 == 0 {
			imgrot = imgrot.flipvert().rotate90()
		}
		for y := 0; y < len(img)-ysize; y++ {
			for x := 0; x < len(img[0])-xsize; x++ {
				match := true
				for _, offset := range offsets {
					if !imgrot[y+offset[1]][x+offset[0]] {
						match = false
						break
					}
				}
				if match {
					monsterCount++
				}
			}
		}
	}
	return monsterCount
}
