package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteWarehouseAppRequest Request Object
type BatchDeleteWarehouseAppRequest struct {
	Body *BatchDeleteWarehouseAppReq `json:"body,omitempty"`
}

func (o BatchDeleteWarehouseAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteWarehouseAppRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteWarehouseAppRequest", string(data)}, " ")
}
