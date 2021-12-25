package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateProductTopicResponse struct {
	// 归属产品ID

	ProductId *int32 `json:"product_id,omitempty"`
	// 产品主题ID

	TopicId *string `json:"topic_id,omitempty"`
	// 主题权限 0-发布 1-订阅

	Permission *int32 `json:"permission,omitempty"`
	// 主题名称

	TopicName *string `json:"topic_name,omitempty"`
	// 版本号

	Version *string `json:"version,omitempty"`
	// 描述

	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateProductTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductTopicResponse struct{}"
	}

	return strings.Join([]string{"CreateProductTopicResponse", string(data)}, " ")
}
