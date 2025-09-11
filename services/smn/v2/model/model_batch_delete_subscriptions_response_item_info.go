package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteSubscriptionsResponseItemInfo struct {

	// 请求的唯一标识ID。
	RequestId string `json:"request_id"`

	// 返回码。
	HttpCode int32 `json:"http_code"`

	// 服务异常错误信息编码。
	Code *string `json:"code,omitempty"`

	// 服务异常错误信息描述。
	Message *string `json:"message,omitempty"`
}

func (o BatchDeleteSubscriptionsResponseItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSubscriptionsResponseItemInfo struct{}"
	}

	return strings.Join([]string{"BatchDeleteSubscriptionsResponseItemInfo", string(data)}, " ")
}
