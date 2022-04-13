package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LinksItem struct {
	// 备份文件名称。

	FileName *string `json:"file_name,omitempty"`
	// 备份文件下载链接地址。

	Link *string `json:"link,omitempty"`
}

func (o LinksItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinksItem struct{}"
	}

	return strings.Join([]string{"LinksItem", string(data)}, " ")
}
