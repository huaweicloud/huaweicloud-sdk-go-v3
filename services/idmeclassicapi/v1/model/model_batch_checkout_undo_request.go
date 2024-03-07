package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckoutUndoRequest Request Object
type BatchCheckoutUndoRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionUndoCheckOutDto `json:"body,omitempty"`
}

func (o BatchCheckoutUndoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckoutUndoRequest struct{}"
	}

	return strings.Join([]string{"BatchCheckoutUndoRequest", string(data)}, " ")
}
