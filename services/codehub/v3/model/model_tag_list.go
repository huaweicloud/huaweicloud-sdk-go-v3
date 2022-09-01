package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagList struct {

	// 指定仓库的标签列表
	Tags *[]Tag `json:"tags,omitempty" xml:"tags"`

	// 指定仓库的标签总数
	Total *int32 `json:"total,omitempty" xml:"total"`
}

func (o TagList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagList struct{}"
	}

	return strings.Join([]string{"TagList", string(data)}, " ")
}
