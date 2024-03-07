package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckoutUndoByAdminRequest Request Object
type BatchCheckoutUndoByAdminRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionUndoCheckOutDto `json:"body,omitempty"`
}

func (o BatchCheckoutUndoByAdminRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckoutUndoByAdminRequest struct{}"
	}

	return strings.Join([]string{"BatchCheckoutUndoByAdminRequest", string(data)}, " ")
}
