package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceTag 独享引擎实例的标签信息
type InstanceTag struct {

	// 标签的键
	Key *string `json:"key,omitempty"`

	// 标签的值
	Values *[]string `json:"values,omitempty"`
}

func (o InstanceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceTag struct{}"
	}

	return strings.Join([]string{"InstanceTag", string(data)}, " ")
}
