package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签信息
type TagsInfo struct {

	// 功能说明：键。同一资源的key值不能重复。
	Key *string `json:"key,omitempty" xml:"key"`

	// 功能说明：值列表。
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o TagsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsInfo struct{}"
	}

	return strings.Join([]string{"TagsInfo", string(data)}, " ")
}
