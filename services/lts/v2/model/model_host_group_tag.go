package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostGroupTag 标签信息
type HostGroupTag struct {

	// 标签Key
	Key *string `json:"key,omitempty"`

	// 标签Value
	Value *string `json:"value,omitempty"`
}

func (o HostGroupTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostGroupTag struct{}"
	}

	return strings.Join([]string{"HostGroupTag", string(data)}, " ")
}
