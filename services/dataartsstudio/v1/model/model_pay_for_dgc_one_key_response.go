package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PayForDgcOneKeyResponse Response Object
type PayForDgcOneKeyResponse struct {

	// 订单Id
	OrderId *string `json:"order_id,omitempty"`

	// 实例Id
	ResourceId     *string `json:"resource_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PayForDgcOneKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PayForDgcOneKeyResponse struct{}"
	}

	return strings.Join([]string{"PayForDgcOneKeyResponse", string(data)}, " ")
}
