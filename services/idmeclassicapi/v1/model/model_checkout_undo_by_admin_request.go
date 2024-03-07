package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckoutUndoByAdminRequest Request Object
type CheckoutUndoByAdminRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionUndoCheckOutDto `json:"body,omitempty"`
}

func (o CheckoutUndoByAdminRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckoutUndoByAdminRequest struct{}"
	}

	return strings.Join([]string{"CheckoutUndoByAdminRequest", string(data)}, " ")
}
