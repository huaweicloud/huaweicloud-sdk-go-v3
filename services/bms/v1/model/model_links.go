package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// links字段数据结构说明
type Links struct {

	// 快捷链接标记名称
	Rel *string `json:"rel,omitempty" xml:"rel"`

	// 对应快捷链接
	Href *string `json:"href,omitempty" xml:"href"`
}

func (o Links) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Links struct{}"
	}

	return strings.Join([]string{"Links", string(data)}, " ")
}
