package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchShowPhoneConnectInfosResponse Response Object
type BatchShowPhoneConnectInfosResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 云手机接入信息列表
	ConnectInfos *[]ConnectInfo `json:"connect_infos,omitempty"`

	// 错误信息
	Errors         *[]ConnectErrorInfo `json:"errors,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o BatchShowPhoneConnectInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowPhoneConnectInfosResponse struct{}"
	}

	return strings.Join([]string{"BatchShowPhoneConnectInfosResponse", string(data)}, " ")
}
