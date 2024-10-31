package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceTag struct {

	// 标签键
	Key *string `json:"key,omitempty"`

	// 标签值
	Value *string `json:"value,omitempty"`

	// 标签更新时间
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
