package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCloudPhoneServerModelResponse Response Object
type ChangeCloudPhoneServerModelResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 订单ID，不超过64个字节
	OrderId *string `json:"order_id,omitempty"`

	// 产品ID，不超过64个字节
	ProductId      *string `json:"product_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeCloudPhoneServerModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCloudPhoneServerModelResponse struct{}"
	}

	return strings.Join([]string{"ChangeCloudPhoneServerModelResponse", string(data)}, " ")
}
