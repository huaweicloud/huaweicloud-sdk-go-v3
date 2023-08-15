package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenApiCalledRecord OpenApi调用记录
type OpenApiCalledRecord struct {

	// 调用API的user_name
	UserName *string `json:"user_name,omitempty"`

	// 调用API的user_id
	UserId *string `json:"user_id,omitempty"`

	// 调用API的domain_name
	DomainName *string `json:"domain_name,omitempty"`

	// 调用API的domain_id
	DomainId *string `json:"domain_id,omitempty"`

	// 调用API的URL
	RequestUrl *string `json:"request_url,omitempty"`

	// http请求方法
	RequestMethod *string `json:"request_method,omitempty"`

	// http状态码
	ResponseCode *string `json:"response_code,omitempty"`

	// 调用API失败原因
	FailReason *string `json:"fail_reason,omitempty"`

	// 调用API的时间（Unix timestamp），单位：毫秒
	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o OpenApiCalledRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenApiCalledRecord struct{}"
	}

	return strings.Join([]string{"OpenApiCalledRecord", string(data)}, " ")
}
