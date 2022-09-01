package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ConfirmConsumptionMessagesResponse struct {

	// 确认成功的数目（如果为N，则表示前N条消息确认成功）。
	Success *int32 `json:"success,omitempty" xml:"success"`

	// 确认失败的数目（如果为N，则表示后N条消息确认失败）。
	Fail           *int32 `json:"fail,omitempty" xml:"fail"`
	HttpStatusCode int    `json:"-"`
}

func (o ConfirmConsumptionMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmConsumptionMessagesResponse struct{}"
	}

	return strings.Join([]string{"ConfirmConsumptionMessagesResponse", string(data)}, " ")
}
