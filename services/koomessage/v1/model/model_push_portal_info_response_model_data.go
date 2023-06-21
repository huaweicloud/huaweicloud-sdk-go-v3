package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 消息体。
type PushPortalInfoResponseModelData struct {

	// 主页申请记录ID。
	LogId *string `json:"log_id,omitempty"`

	// 返回信息。
	Message *string `json:"message,omitempty"`
}

func (o PushPortalInfoResponseModelData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushPortalInfoResponseModelData struct{}"
	}

	return strings.Join([]string{"PushPortalInfoResponseModelData", string(data)}, " ")
}
