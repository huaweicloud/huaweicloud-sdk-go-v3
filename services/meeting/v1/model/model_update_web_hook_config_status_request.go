package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateWebHookConfigStatusRequest struct {

	// 订阅配置记录ID。
	Id string `json:"id"`

	// 事件推送状态。 * 0：启用 * 1：禁用 * 2：锁定
	Status int32 `json:"status"`
}

func (o UpdateWebHookConfigStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWebHookConfigStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateWebHookConfigStatusRequest", string(data)}, " ")
}
