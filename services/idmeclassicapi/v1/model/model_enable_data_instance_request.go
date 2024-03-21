package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDataInstanceRequest Request Object
type EnableDataInstanceRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistObjectIdsModifierDto `json:"body,omitempty"`
}

func (o EnableDataInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDataInstanceRequest struct{}"
	}

	return strings.Join([]string{"EnableDataInstanceRequest", string(data)}, " ")
}
