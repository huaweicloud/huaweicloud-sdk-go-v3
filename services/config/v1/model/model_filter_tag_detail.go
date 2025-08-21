package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FilterTagDetail 标签对象
type FilterTagDetail struct {

	// 标签key
	Key *string `json:"key,omitempty"`

	// 标签值列表
	Values *[]string `json:"values,omitempty"`
}

func (o FilterTagDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilterTagDetail struct{}"
	}

	return strings.Join([]string{"FilterTagDetail", string(data)}, " ")
}
