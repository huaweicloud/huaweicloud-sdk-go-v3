package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DebugApiResponse Response Object
type DebugApiResponse struct {

	// 请求编号
	RequestId *string `json:"request_id,omitempty"`

	// 请求url
	Url *string `json:"url,omitempty"`

	// 调试结果
	Result *string `json:"result,omitempty"`

	// 调试耗时
	Timeout *int64 `json:"timeout,omitempty"`

	// 是否调试成功
	Success *bool `json:"success,omitempty"`

	RequestHeader *ApiTestRequestHeader `json:"request_header,omitempty"`

	ResponseHeader *ApiTestResponseHeader `json:"response_header,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o DebugApiResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DebugApiResponse struct{}"
	}

	return strings.Join([]string{"DebugApiResponse", string(data)}, " ")
}
