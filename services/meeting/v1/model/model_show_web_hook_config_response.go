package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWebHookConfigResponse Response Object
type ShowWebHookConfigResponse struct {

	// 结果码。
	ReturnCode int32 `json:"returnCode"`

	// 结果描述。
	ReturnDesc *string `json:"returnDesc,omitempty"`

	// 订阅配置记录ID。
	Id *string `json:"id,omitempty"`

	// 订阅ID。
	SubscriberId *string `json:"subscriberId,omitempty"`

	// 订阅url。
	Url *string `json:"url,omitempty"`

	// 事件推送状态。 * 0：已启用 * 1：未启动 * 2：已锁定
	Status         *int32 `json:"status,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowWebHookConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWebHookConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowWebHookConfigResponse", string(data)}, " ")
}
