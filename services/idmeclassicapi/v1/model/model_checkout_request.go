package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckoutRequest Request Object
type CheckoutRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionCheckOutDto `json:"body,omitempty"`
}

func (o CheckoutRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckoutRequest struct{}"
	}

	return strings.Join([]string{"CheckoutRequest", string(data)}, " ")
}
