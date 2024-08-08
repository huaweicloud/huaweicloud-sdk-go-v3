package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWarehouseAppRequest Request Object
type CreateWarehouseAppRequest struct {
	Body *CreateWarehouseAppReq `json:"body,omitempty"`
}

func (o CreateWarehouseAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWarehouseAppRequest struct{}"
	}

	return strings.Join([]string{"CreateWarehouseAppRequest", string(data)}, " ")
}
