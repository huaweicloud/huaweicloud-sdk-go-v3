package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWarehouseAppResponse Response Object
type DeleteWarehouseAppResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWarehouseAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWarehouseAppResponse struct{}"
	}

	return strings.Join([]string{"DeleteWarehouseAppResponse", string(data)}, " ")
}
