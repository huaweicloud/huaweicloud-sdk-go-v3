package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModelRequest Request Object
type UpdateModelRequest struct {

	// 供应商id。
	ProviderId string `json:"provider_id"`

	// 模型id。
	ModelId string `json:"model_id"`

	Body *UpdateModelReq `json:"body,omitempty"`
}

func (o UpdateModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModelRequest struct{}"
	}

	return strings.Join([]string{"UpdateModelRequest", string(data)}, " ")
}
