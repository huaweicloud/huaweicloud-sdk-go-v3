package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LinksItem struct {
	// 对应快捷链接。

	Href string `json:"href"`
	// 快捷链接标记名称。

	Rel string `json:"rel"`
}

func (o LinksItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinksItem struct{}"
	}

	return strings.Join([]string{"LinksItem", string(data)}, " ")
}
