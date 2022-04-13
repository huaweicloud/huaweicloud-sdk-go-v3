package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteWebHookConfigRequest struct {
	// 订阅配置记录id

	Id string `json:"id"`
}

func (o DeleteWebHookConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWebHookConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteWebHookConfigRequest", string(data)}, " ")
}
