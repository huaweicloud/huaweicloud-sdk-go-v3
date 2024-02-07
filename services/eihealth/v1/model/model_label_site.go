package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LabelSite 靶点分子设计标记位点
type LabelSite struct {

	// 索引
	Index *[]int32 `json:"index,omitempty"`

	// 标记位点名称
	Name *[]string `json:"name,omitempty"`

	// 位点三维坐标集
	Coordinate *[][]float32 `json:"coordinate,omitempty"`
}

func (o LabelSite) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelSite struct{}"
	}

	return strings.Join([]string{"LabelSite", string(data)}, " ")
}
