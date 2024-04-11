package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RefreshRequest Request Object
type RefreshRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoPersistObjectIdsDto `json:"body,omitempty"`
}

func (o RefreshRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefreshRequest struct{}"
	}

	return strings.Join([]string{"RefreshRequest", string(data)}, " ")
}
