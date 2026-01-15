package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTagInfo 标签信息体
type ResourceTagInfo struct {

	// 标签键。例如键值对{“aaa”:\"bbb\"}的key为\"aaa\"
	Key *string `json:"key,omitempty"`

	// 标签值。例如键值对{“aaa”:\"bbb\"}的value为\"bbb\"
	Value *string `json:"value,omitempty"`
}

func (o ResourceTagInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTagInfo struct{}"
	}

	return strings.Join([]string{"ResourceTagInfo", string(data)}, " ")
}
