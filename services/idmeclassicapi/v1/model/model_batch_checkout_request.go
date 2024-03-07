package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckoutRequest Request Object
type BatchCheckoutRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListVersionModelVersionCheckOutDto `json:"body,omitempty"`
}

func (o BatchCheckoutRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckoutRequest struct{}"
	}

	return strings.Join([]string{"BatchCheckoutRequest", string(data)}, " ")
}
