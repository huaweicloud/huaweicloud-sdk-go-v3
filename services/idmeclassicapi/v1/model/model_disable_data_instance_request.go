package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableDataInstanceRequest Request Object
type DisableDataInstanceRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistObjectIdsModifierDto `json:"body,omitempty"`
}

func (o DisableDataInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableDataInstanceRequest struct{}"
	}

	return strings.Join([]string{"DisableDataInstanceRequest", string(data)}, " ")
}
