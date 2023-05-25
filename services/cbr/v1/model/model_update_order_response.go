package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateOrderResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrderResponse struct{}"
	}

	return strings.Join([]string{"UpdateOrderResponse", string(data)}, " ")
}
