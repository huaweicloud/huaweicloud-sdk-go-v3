package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签对象列表。
type TagList struct {
	Tag *TagReq `json:"tag,omitempty"`
}

func (o TagList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagList struct{}"
	}

	return strings.Join([]string{"TagList", string(data)}, " ")
}
