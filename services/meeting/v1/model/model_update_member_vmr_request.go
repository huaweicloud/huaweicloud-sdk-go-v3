package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateMemberVmrRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN, 英文为en-US
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 云会议室唯一标识。
	Id string `json:"id"`

	Body *ModVmrDto `json:"body,omitempty"`
}

func (o UpdateMemberVmrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberVmrRequest struct{}"
	}

	return strings.Join([]string{"UpdateMemberVmrRequest", string(data)}, " ")
}
