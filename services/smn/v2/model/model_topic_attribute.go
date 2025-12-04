package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopicAttribute struct {

	// topic的访问策略
	AccessPolicy *string `json:"access_policy,omitempty"`

	// topic的访问策略创建时间。时间格式为UTC时间，YYYY-MM-DDTHH:MM:SSZ。
	CreateTime *string `json:"create_time,omitempty"`

	// topic的访问策略更新时间。时间格式为UTC时间，YYYY-MM-DDTHH:MM:SSZ。
	UpdateTime *string `json:"update_time,omitempty"`

	// topic的简介
	Introduction *string `json:"introduction,omitempty"`
}

func (o TopicAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicAttribute struct{}"
	}

	return strings.Join([]string{"TopicAttribute", string(data)}, " ")
}
