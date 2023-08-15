package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushMenuInfoResponseModelData 消息体。
type PushMenuInfoResponseModelData struct {

	// 菜单申请记录ID。
	LogId *string `json:"log_id,omitempty"`

	// 返回信息。
	Message *string `json:"message,omitempty"`
}

func (o PushMenuInfoResponseModelData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushMenuInfoResponseModelData struct{}"
	}

	return strings.Join([]string{"PushMenuInfoResponseModelData", string(data)}, " ")
}
