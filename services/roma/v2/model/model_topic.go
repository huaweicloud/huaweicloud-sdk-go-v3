package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Topic struct {
	// TOPIC的ID

	Id *int32 `json:"id,omitempty"`
	// TOPIC的名称

	Name *string `json:"name,omitempty"`
	// TOPIC描述

	Description *string `json:"description,omitempty"`
	// TOPIC权限, 主题权限 0-发布 1-订阅

	Permission *int32 `json:"permission,omitempty"`
	// TOPIC类型 0-基础TOPIC 1-用户自定义TOPIC

	IsPrivate *int32 `json:"is_private,omitempty"`
}

func (o Topic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Topic struct{}"
	}

	return strings.Join([]string{"Topic", string(data)}, " ")
}
