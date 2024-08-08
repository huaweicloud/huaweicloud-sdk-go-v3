package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateModelReq struct {

	// 模型名称，取值范围：[5,32]，允许大小写字母、数字、下划线(_)、中划线(-)和空格,只能以字母开头
	Name string `json:"name"`

	// 模型描述信息
	Description *string `json:"description,omitempty"`

	Type *ModelType `json:"type"`

	File *ModelFile `json:"file"`

	// 是否打开组织共享
	Shareable *bool `json:"shareable,omitempty"`

	// 基模型id
	BaseModelId *string `json:"base_model_id,omitempty"`
}

func (o CreateModelReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelReq struct{}"
	}

	return strings.Join([]string{"CreateModelReq", string(data)}, " ")
}
