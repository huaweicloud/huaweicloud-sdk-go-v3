package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoDto struct {
	// 页码

	Offset *int32 `json:"offset,omitempty"`
	// 每页记录数

	Limit *int32 `json:"limit,omitempty"`
}

func (o PageInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoDto struct{}"
	}

	return strings.Join([]string{"PageInfoDto", string(data)}, " ")
}
