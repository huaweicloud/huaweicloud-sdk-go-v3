package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Tag struct {

	// 标签键。
	Key *string `json:"key,omitempty"`

	// 标签值集合。  约束：最多包含10个value。标签列表中的标签value值不允许重复。标签列表如果为空列表，表示匹配任意值。标签列表中多个value之间是“或”的关系，在key已经满足要求的前提下，满足请求中的某个value就会匹配出来。
	Values *[]string `json:"values,omitempty"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
