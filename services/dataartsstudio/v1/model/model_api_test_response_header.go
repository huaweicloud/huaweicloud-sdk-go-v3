package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiTestResponseHeader struct {

	// 是否成功
	ResultStatus *string `json:"result_status,omitempty"`

	// 内容大小
	ContentLength *int32 `json:"content_length,omitempty"`

	// 连接状态
	Connection *string `json:"connection,omitempty"`

	// 缓存控制（固定值）
	CacheControl *string `json:"cache_control,omitempty"`

	// 内容类型 （固定值）
	ContentType *string `json:"content_type,omitempty"`

	// 日期
	Date *string `json:"date,omitempty"`

	// 请求ID
	XRequestId *string `json:"x_request_id,omitempty"`
}

func (o ApiTestResponseHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiTestResponseHeader struct{}"
	}

	return strings.Join([]string{"ApiTestResponseHeader", string(data)}, " ")
}
