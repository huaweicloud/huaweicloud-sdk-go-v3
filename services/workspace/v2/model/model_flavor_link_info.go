package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlavorLinkInfo struct {

	// 快捷链接标记名称。
	Rel *string `json:"rel,omitempty"`

	// 对应快捷链接。
	Hrel *string `json:"hrel,omitempty"`
}

func (o FlavorLinkInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorLinkInfo struct{}"
	}

	return strings.Join([]string{"FlavorLinkInfo", string(data)}, " ")
}
