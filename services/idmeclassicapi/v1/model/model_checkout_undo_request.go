package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckoutUndoRequest Request Object
type CheckoutUndoRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionUndoCheckOutDto `json:"body,omitempty"`
}

func (o CheckoutUndoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckoutUndoRequest struct{}"
	}

	return strings.Join([]string{"CheckoutUndoRequest", string(data)}, " ")
}
