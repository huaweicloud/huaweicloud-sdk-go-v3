package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Location struct {

	// 节点在作业画布上的横轴位置
	X string `json:"x"`

	// 节点在作业画布上的纵轴位置
	Y string `json:"y"`
}

func (o Location) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Location struct{}"
	}

	return strings.Join([]string{"Location", string(data)}, " ")
}
