package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckoutAndUpdateRequest Request Object
type CheckoutAndUpdateRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoVersionModelVersionCheckoutAndUpdateDtoVersionModel `json:"body,omitempty"`
}

func (o CheckoutAndUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckoutAndUpdateRequest struct{}"
	}

	return strings.Join([]string{"CheckoutAndUpdateRequest", string(data)}, " ")
}
