package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteWarehouseAppResponse Response Object
type BatchDeleteWarehouseAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteWarehouseAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteWarehouseAppResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteWarehouseAppResponse", string(data)}, " ")
}
