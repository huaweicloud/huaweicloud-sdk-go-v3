package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWebHookConfigRequest Request Object
type DeleteWebHookConfigRequest struct {

	// 订阅配置记录ID。
	Id string `json:"id"`
}

func (o DeleteWebHookConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWebHookConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteWebHookConfigRequest", string(data)}, " ")
}
