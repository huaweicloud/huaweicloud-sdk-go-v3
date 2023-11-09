package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tags 标签
type Tags struct {

	// 标签key。
	Key *string `json:"key,omitempty"`

	// 标签value。
	Values *[]string `json:"values,omitempty"`
}

func (o Tags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tags struct{}"
	}

	return strings.Join([]string{"Tags", string(data)}, " ")
}
