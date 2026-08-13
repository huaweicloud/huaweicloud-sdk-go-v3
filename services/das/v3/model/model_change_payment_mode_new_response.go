package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePaymentModeNewResponse Response Object
type ChangePaymentModeNewResponse struct {

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 可恢复为免费实例的时间
	CanSetFreeTime float32 `json:"can_set_free_time,omitempty"`

	// 错误信息
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangePaymentModeNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePaymentModeNewResponse struct{}"
	}

	return strings.Join([]string{"ChangePaymentModeNewResponse", string(data)}, " ")
}
