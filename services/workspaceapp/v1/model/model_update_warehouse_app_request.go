package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWarehouseAppRequest Request Object
type UpdateWarehouseAppRequest struct {

	// 应用仓库中的应用记录ID。
	Id string `json:"id"`

	Body *UpdateWarehouseAppReq `json:"body,omitempty"`
}

func (o UpdateWarehouseAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWarehouseAppRequest struct{}"
	}

	return strings.Join([]string{"UpdateWarehouseAppRequest", string(data)}, " ")
}
