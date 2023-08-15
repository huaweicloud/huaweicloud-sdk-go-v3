package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VertexLocationDto struct {

	// 子任务的画布横坐标
	X *float64 `json:"x,omitempty"`

	// 子任务的画布纵坐标
	Y *float64 `json:"y,omitempty"`
}

func (o VertexLocationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VertexLocationDto struct{}"
	}

	return strings.Join([]string{"VertexLocationDto", string(data)}, " ")
}
