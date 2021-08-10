package model

import (
	"encoding/json"

	"strings"
)

type Point struct {
	// 坐标横轴数值。

	X float64 `json:"x"`
	// 坐标纵轴数值。

	Y float64 `json:"y"`
}

func (o Point) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Point struct{}"
	}

	return strings.Join([]string{"Point", string(data)}, " ")
}
