package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlavorLink struct {

	// 快捷链接标记名称。
	Rel *string `json:"rel,omitempty"`

	// 对应快捷链接。
	Hrel *string `json:"hrel,omitempty"`
}

func (o FlavorLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorLink struct{}"
	}

	return strings.Join([]string{"FlavorLink", string(data)}, " ")
}
