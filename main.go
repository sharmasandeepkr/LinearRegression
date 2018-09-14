package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	//"image/color"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	//"github.com/gonum/stat/distuv"
)

type xy struct {
	x float64
	y float64
}

func main() {

	var xys3 []xy
	// var xys4 []xy
	// var xys5 []xy

	f3, err := os.Open("./data3.txt")
	if err != nil {
		fmt.Errorf("unable to open file %v", err)
	}
	// f4, err := os.Open("./data4.txt")
	// if err != nil {
	// 	fmt.Errorf("unable to open one of file%v", err)
	// }
	// f5, err := os.Open("./data5.txt")
	// if err != nil {
	// 	fmt.Errorf("unable to open 5th file%v", err)
	// }

	defer f3.Close()
	scanner := bufio.NewScanner(f3)
	for scanner.Scan() {
		xyline := scanner.Text()
		xys3 = append(xys3, returnxy(xyline))
	}

	// defer f4.Close()
	// scanner = bufio.NewScanner(f4)
	// for scanner.Scan() {
	// 	xyline := scanner.Text()
	// 	xys4 = append(xys4, returnxy(xyline))
	// }

	// defer f5.Close()
	// scanner = bufio.NewScanner(f5)
	// for scanner.Scan() {
	// 	xyline := scanner.Text()
	// 	xys5 = append(xys5, returnxy(xyline))
	// }
	//fmt.Println(xys)
	plotgraph(xys3 /*, xys4, xys5*/)
	if err = scanner.Err(); err != nil {
		fmt.Errorf("uable to scan %v", err)
	}

}

func plotgraph(xys3 /*, xys4, xys5 */ []xy) {
	p, err := plot.New()
	if err != nil {
		fmt.Errorf("unable to plot from given xys%v", err)
	}
	p.Title.Text = "scattered graph"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddScatters(p, rxys(xys3))
	if err != nil {
		fmt.Errorf("unable to scatter%v", err)
	}

	m1, c1 := getiline(rxys(xys3))
	// m2, c2 := getiline(rxys(xys4))
	// m3, c3 := getiline(rxys(xys5))

	err = plotutil.AddLinePoints(p, putiline(xys3, m1, c1)) // "FourthFile", putiline(xys4, m2, c2),
	// "fifthFile", putiline(xys5, m3, c3),

	if err != nil {
		fmt.Errorf("unable to make line%v", err)
	}
	if err = p.Save(4*vg.Inch, 4*vg.Inch, "out.png"); err != nil {
		fmt.Errorf("unable to save out.png bcz %v", err)
	}
}

func getiline(xys plotter.XYs) (float64, float64) {
	if len(xys) < 1 {
		fmt.Errorf("your xys are empty ")
	}
	var x []float64
	// var x [] plotter.XYs.X
	for k := 0; k < len(xys); k++ {
		x = append(x, xys[k].X)
	}
	// for v := range xys {

	// }
	minc := 6.0
	minm := 3.0
	// y=m*x+c
	mincost := 999.999
	for m := 0.3; m < 1; m += 0.01 {
		for c := 0.0; c < 6; c += 0.1 {
			for _ = range x {
				cost := computeCost(m, c, x, xys)
				if cost < mincost {
					mincost = cost
					minc = c
					minm = m

				}
			}

		}
	}
	fmt.Println("mincost= ", mincost, "minc= ", minc, "minm= ", minm)
	return minm, minc
}

func computeCost(m float64, c float64, i []float64, xys plotter.XYs) float64 {

	if len(i) == 1 || len(xys) == 1 {
		return math.Sqrt(math.Pow((i[0]-xys[0].X), 2) + math.Pow((m*i[0]+c)-xys[0].Y, 2))
	}
	var totalCost float64
	var slicedX []float64
	var nxys plotter.XYs
	for _ = range i {
		slicedX = i[1:]
		nxys = xys[1:]
		totalCost = lineCost(i[0], m*i[0]+c, xys[0].X, xys[0].Y) + computeCost(m, c, slicedX, nxys)
	}
	fmt.Println(totalCost, "totalcost")
	return totalCost

}

func lineCost(xp float64, yp float64, xysx float64, xysy float64) float64 {
	return math.Sqrt(math.Pow((xp-xysx), 2) + math.Pow(yp-xysy, 2))
}
func putiline(xys []xy, m float64, c float64) plotter.XYs {
	pts := make(plotter.XYs, len(xys))
	i := 0
	for v := range pts {

		pts[v].X = xys[i].x
		pts[v].Y = xys[i].x*m + c
		i++
	}
	return pts
}

func rxys(xys []xy) plotter.XYs {
	pts := make(plotter.XYs, len(xys))
	i := 0
	for v := range pts {

		pts[v].X = xys[i].x
		pts[v].Y = xys[i].y
		i++
	}
	return pts
}

func returnxy(xyline string) xy {
	var xy xy
	var err error
	xys := strings.Split(xyline, ",")
	xy.x, err = strconv.ParseFloat(xys[0], 64)
	if err != nil {
		fmt.Errorf("cannot convert this string to float %v", err)

	}
	xy.y, err = strconv.ParseFloat(xys[1], 64)
	if err != nil {
		fmt.Errorf("unable to convert this y to %v", err)
	}
	return xy

}
