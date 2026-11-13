package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TbPosInfo struct {

	// 原始名称
	OriginName *string `json:"origin_name,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 开始
	Start *int32 `json:"start,omitempty"`

	// 结束
	End *int32 `json:"end,omitempty"`
}

func (o TbPosInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TbPosInfo struct{}"
	}

	return strings.Join([]string{"TbPosInfo", string(data)}, " ")
}
