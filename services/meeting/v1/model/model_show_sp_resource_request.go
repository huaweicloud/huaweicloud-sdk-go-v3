package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSpResourceRequest Request Object
type ShowSpResourceRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN，英文为en-US。
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 是否在查询资源信息时返回SP自主规划的媒体接入分组id，默认不查询。
	QueryGroup *bool `json:"queryGroup,omitempty"`
}

func (o ShowSpResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpResourceRequest struct{}"
	}

	return strings.Join([]string{"ShowSpResourceRequest", string(data)}, " ")
}
