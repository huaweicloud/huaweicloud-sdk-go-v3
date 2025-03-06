package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCloudPhoneSingleServerResponse Response Object
type CreateCloudPhoneSingleServerResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 订单ID，不超过64个字节。
	OrderId *string `json:"order_id,omitempty"`

	// 产品ID，不超过64个字节。
	ProductId *string `json:"product_id,omitempty"`

	// 服务器ID列表。
	ServerIds *[]string `json:"server_ids,omitempty"`

	// 任务错误码说明。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 任务错误码。
	ErrorCode      *string `json:"error_code,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCloudPhoneSingleServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudPhoneSingleServerResponse struct{}"
	}

	return strings.Join([]string{"CreateCloudPhoneSingleServerResponse", string(data)}, " ")
}
