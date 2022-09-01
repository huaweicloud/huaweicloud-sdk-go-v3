package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCloudPhoneServerResponse struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id" xml:"request_id"`

	// 订单ID，不超过64个字节
	OrderId string `json:"order_id" xml:"order_id"`

	// 产品ID，不超过64个字节
	ProductId      string `json:"product_id" xml:"product_id"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateCloudPhoneServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudPhoneServerResponse struct{}"
	}

	return strings.Join([]string{"CreateCloudPhoneServerResponse", string(data)}, " ")
}
