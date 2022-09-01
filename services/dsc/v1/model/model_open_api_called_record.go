package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenApi调用记录
type OpenApiCalledRecord struct {

	// 调用API的user_name
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 调用API的user_id
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 调用API的domain_name
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// 调用API的domain_id
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 调用API的URL
	RequestUrl *string `json:"request_url,omitempty" xml:"request_url"`

	// http请求方法
	RequestMethod *string `json:"request_method,omitempty" xml:"request_method"`

	// http状态码
	ResponseCode *string `json:"response_code,omitempty" xml:"response_code"`

	// 调用API失败原因
	FailReason *string `json:"fail_reason,omitempty" xml:"fail_reason"`

	// 调用API的时间（Unix timestamp），单位：毫秒
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp"`
}

func (o OpenApiCalledRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenApiCalledRecord struct{}"
	}

	return strings.Join([]string{"OpenApiCalledRecord", string(data)}, " ")
}
