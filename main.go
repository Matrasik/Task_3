package main

import (
	"Beautiful_Code/distance"
	"Beautiful_Code/distance/point"
	"Beautiful_Code/navigator"
	"fmt"
)

func main() {
	//line := distance.NewLine2d(*point.NewPoint2d(1, 0), *point.NewPoint2d(0, 1))
	//
	//fmt.Println(line)
	//
	//distance2dLine := line.Distance()
	//fmt.Printf("line 2d %T имеет расстояние %v\n", line, distance2dLine)
	//
	//line3d1 := distance.NewLine3d(*point.NewPoint3d(1, 0, 1), *point.NewPoint3d(0, 1, 0))
	//
	//fmt.Println(line3d1)
	//
	//line3dDistance, _ := line3d1.Distance()
	//fmt.Printf("Line3d имеет расстояние %v\n", line3dDistance)
	//
	//path3d1 := distance.NewPath3d(make([]point.Point3d, 0))
	//path3d1.AddPoint(*point.NewPoint3d(1, 1, 1)).AddPoint(*point.NewPoint3d(2, 3, 3)).AddPoint(*point.NewPoint3d(2, 3, 4))
	//
	//distancePath3d, _ := path3d1.Distance()
	//fmt.Printf("path3d полный путь %v\n", distancePath3d)
	//
	lineLatLon := distance.NewLineLatLon(*point.NewPoinLatLon(55.7539, 37.6208), *point.NewPoinLatLon(59.9398, 30.3146))
	distanceLatLon, _ := lineLatLon.Distance()
	fmt.Println(distanceLatLon)

	navi := navigator.NewNavigator([]navigator.AllDistance{lineLatLon})

	fullPath, err := navi.Path()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Полный путь по навигатору %v\n", fullPath)

}
