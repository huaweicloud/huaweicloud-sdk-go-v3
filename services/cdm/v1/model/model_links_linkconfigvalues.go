package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 连接参数配置，请参见link-config-values参数说明
type LinksLinkconfigvalues struct {
	// 连接配置参数数据结构，请参见configs参数说明。

	Configs *[]Configs `json:"configs,omitempty"`
}

func (o LinksLinkconfigvalues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinksLinkconfigvalues struct{}"
	}

	return strings.Join([]string{"LinksLinkconfigvalues", string(data)}, " ")
}
