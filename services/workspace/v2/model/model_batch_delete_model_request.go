package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteModelRequest Request Object
type BatchDeleteModelRequest struct {

	// 供应商id。
	ProviderId string `json:"provider_id"`

	Body *ModelBatchDeleteReq `json:"body,omitempty"`
}

func (o BatchDeleteModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteModelRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteModelRequest", string(data)}, " ")
}
