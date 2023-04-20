package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Position struct {

	// 像素坐标x
	X int32 `json:"x"`

	// 像素坐标x
	Y int32 `json:"y"`
}

func (o Position) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Position struct{}"
	}

	return strings.Join([]string{"Position", string(data)}, " ")
}
