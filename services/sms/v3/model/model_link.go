package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Link struct {

	// API的url地址。
	Href *string `json:"href,omitempty"`

	// 取值为“self”，表示href为本地链接。
	Rel *string `json:"rel,omitempty"`
}

func (o Link) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Link struct{}"
	}

	return strings.Join([]string{"Link", string(data)}, " ")
}
