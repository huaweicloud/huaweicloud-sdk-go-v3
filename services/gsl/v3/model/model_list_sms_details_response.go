package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmsDetailsResponse Response Object
type ListSmsDetailsResponse struct {

	// 每页的记录数
	Limit *int64 `json:"limit,omitempty"`

	// 页码，最小值是1，最大值为1000000。默认值是1.
	Offset *int64 `json:"offset,omitempty"`

	// 记录总数
	Count *int64 `json:"count,omitempty"`

	// 短信发送详情列表
	SmsDetails     *[]SmsSendDetailQueryVo `json:"sms_details,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListSmsDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmsDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListSmsDetailsResponse", string(data)}, " ")
}
