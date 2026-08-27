package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModelRequest Request Object
type ShowModelRequest struct {

	// 供应商id。
	ProviderId string `json:"provider_id"`

	// 模型id。
	ModelId string `json:"model_id"`
}

func (o ShowModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModelRequest struct{}"
	}

	return strings.Join([]string{"ShowModelRequest", string(data)}, " ")
}
