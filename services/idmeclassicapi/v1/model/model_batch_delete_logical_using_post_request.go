package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteLogicalUsingPostRequest Request Object
type BatchDeleteLogicalUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistObjectIdsModifierDto `json:"body,omitempty"`
}

func (o BatchDeleteLogicalUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLogicalUsingPostRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteLogicalUsingPostRequest", string(data)}, " ")
}
