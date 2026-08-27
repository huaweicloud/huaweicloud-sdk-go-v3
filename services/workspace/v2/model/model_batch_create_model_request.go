package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateModelRequest Request Object
type BatchCreateModelRequest struct {

	// 供应商id。
	ProviderId string `json:"provider_id"`

	Body *ModelBatchCreateReq `json:"body,omitempty"`
}

func (o BatchCreateModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateModelRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateModelRequest", string(data)}, " ")
}
