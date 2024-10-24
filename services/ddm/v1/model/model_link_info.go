package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LinkInfo struct {

	// 对应该API的URL
	Href *string `json:"href,omitempty"`

	// 取值为“self”，表示href为本地链接。
	Rel *string `json:"rel,omitempty"`
}

func (o LinkInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinkInfo struct{}"
	}

	return strings.Join([]string{"LinkInfo", string(data)}, " ")
}
