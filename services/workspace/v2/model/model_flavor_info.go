package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlavorInfo struct {

	// 桌面对应的规格ID。
	Id *string `json:"id,omitempty"`

	// 桌面对应规格的相关标记快捷链接信息。
	Links *[]FlavorLinkInfo `json:"links,omitempty"`
}

func (o FlavorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorInfo struct{}"
	}

	return strings.Join([]string{"FlavorInfo", string(data)}, " ")
}
