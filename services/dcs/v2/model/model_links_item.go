package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LinksItem struct {

	// 备份文件名称。
	FileName *string `json:"file_name,omitempty" xml:"file_name"`

	// 备份文件下载链接地址。
	Link *string `json:"link,omitempty" xml:"link"`
}

func (o LinksItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LinksItem struct{}"
	}

	return strings.Join([]string{"LinksItem", string(data)}, " ")
}
