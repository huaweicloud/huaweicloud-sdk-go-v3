package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签
type ResourceTagResp struct {

	// 键。同一资源的key值不能重复。
	Key *string `json:"key,omitempty" xml:"key"`

	// 值列表。
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o ResourceTagResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTagResp struct{}"
	}

	return strings.Join([]string{"ResourceTagResp", string(data)}, " ")
}
