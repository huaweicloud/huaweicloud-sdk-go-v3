package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateProductTemplateResponse struct {
	// 产品模板ID

	Id *int32 `json:"id,omitempty"`
	// 产品模板名称

	Name *string `json:"name,omitempty"`
	// 产品模板描述

	Description *string `json:"description,omitempty"`
	// 产品模板状态 0-启用 1-停用

	Status *int32 `json:"status,omitempty"`

	CreatedUser *CreatedUser `json:"created_user,omitempty"`

	LastUpdatedUser *LastUpdatedUser `json:"last_updated_user,omitempty"`
	// 创建时间，timestamp(ms)，使用UTC时区

	CreatedDatetime *int64 `json:"created_datetime,omitempty"`
	// 最后修改时间，timestamp(ms)，使用UTC时区

	LastUpdatedDatetime *int64 `json:"last_updated_datetime,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o CreateProductTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateProductTemplateResponse", string(data)}, " ")
}
