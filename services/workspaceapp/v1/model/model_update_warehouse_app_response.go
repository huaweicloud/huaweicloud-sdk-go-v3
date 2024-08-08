package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWarehouseAppResponse Response Object
type UpdateWarehouseAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateWarehouseAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWarehouseAppResponse struct{}"
	}

	return strings.Join([]string{"UpdateWarehouseAppResponse", string(data)}, " ")
}
