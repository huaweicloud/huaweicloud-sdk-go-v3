package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateWebHookConfigStatusRequest struct {
	// 订阅配置记录id

	Id string `json:"id"`
	// 连接状态： 0表示已启用 ；1表示未启动； 2表示已锁定

	Status int32 `json:"status"`
}

func (o UpdateWebHookConfigStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWebHookConfigStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateWebHookConfigStatusRequest", string(data)}, " ")
}
