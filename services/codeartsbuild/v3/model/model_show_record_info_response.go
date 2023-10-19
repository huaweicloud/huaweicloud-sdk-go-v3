package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRecordInfoResponse Response Object
type ShowRecordInfoResponse struct {

	// 状态
	Success *bool `json:"success,omitempty"`

	// 消息
	Message *string `json:"message,omitempty"`

	// 错误码
	ErrCode *string `json:"err_code,omitempty"`

	Result         *RecordInfoResult `json:"result,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowRecordInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecordInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowRecordInfoResponse", string(data)}, " ")
}
