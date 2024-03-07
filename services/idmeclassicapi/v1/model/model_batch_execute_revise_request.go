package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchExecuteReviseRequest Request Object
type BatchExecuteReviseRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionReviseDto `json:"body,omitempty"`
}

func (o BatchExecuteReviseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExecuteReviseRequest struct{}"
	}

	return strings.Join([]string{"BatchExecuteReviseRequest", string(data)}, " ")
}
