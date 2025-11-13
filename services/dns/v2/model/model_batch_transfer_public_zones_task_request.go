package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchTransferPublicZonesTaskRequest Request Object
type BatchTransferPublicZonesTaskRequest struct {
	Body *BatchTransferPublicZonesTaskRequestBody `json:"body,omitempty"`
}

func (o BatchTransferPublicZonesTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTransferPublicZonesTaskRequest struct{}"
	}

	return strings.Join([]string{"BatchTransferPublicZonesTaskRequest", string(data)}, " ")
}
