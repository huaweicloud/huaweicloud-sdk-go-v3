package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTag 在对接TMS服务的各个接口中使用，面向租户和EPS服务
type ResourceTag struct {

	// 标签的key
	Key string `json:"key"`

	// 标签的value
	Value *string `json:"value,omitempty"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
