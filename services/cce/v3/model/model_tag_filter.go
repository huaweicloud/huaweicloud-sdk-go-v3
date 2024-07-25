package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagFilter struct {

	// 需要过滤的标签key。
	Key *string `json:"key,omitempty"`

	Values *[]string `json:"values,omitempty"`
}

func (o TagFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagFilter struct{}"
	}

	return strings.Join([]string{"TagFilter", string(data)}, " ")
}
