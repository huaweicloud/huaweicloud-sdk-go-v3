package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SysTag struct {

	// 标签键。
	Key *string `json:"key,omitempty"`

	// 标签值。
	Value *string `json:"value,omitempty"`
}

func (o SysTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SysTag struct{}"
	}

	return strings.Join([]string{"SysTag", string(data)}, " ")
}
