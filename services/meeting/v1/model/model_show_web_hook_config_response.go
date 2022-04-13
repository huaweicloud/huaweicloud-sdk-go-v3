package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowWebHookConfigResponse struct {
	// 结果码

	ReturnCode int32 `json:"returnCode"`
	// 结果描述

	ReturnDesc *string `json:"returnDesc,omitempty"`
	// 配置记录id

	Id *string `json:"id,omitempty"`
	// 订阅ID

	SubscriberId *string `json:"subscriberId,omitempty"`
	// 订阅url

	Url *string `json:"url,omitempty"`
	// 连接状态： 0表示已启用 ；1表示未启动； 2表示已锁定

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
