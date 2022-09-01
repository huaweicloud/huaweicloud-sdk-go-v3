package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签信息
type HostGroupTag struct {

	// 标签Key
	Key *string `json:"key,omitempty" xml:"key"`

	// 标签Value
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o HostGroupTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostGroupTag struct{}"
	}

	return strings.Join([]string{"HostGroupTag", string(data)}, " ")
}
