package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Resolution struct {

	// 像素x
	X *int32 `json:"x,omitempty"`

	// 像素y
	Y int32 `json:"y"`
}

func (o Resolution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resolution struct{}"
	}

	return strings.Join([]string{"Resolution", string(data)}, " ")
}
