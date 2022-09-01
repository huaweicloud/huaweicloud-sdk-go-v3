package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowWebHookConfigResponse struct {

	// 结果码
	ReturnCode int32 `json:"returnCode" xml:"returnCode"`

	// 结果描述
	ReturnDesc *string `json:"returnDesc,omitempty" xml:"returnDesc"`

	// 配置记录id
	Id *string `json:"id,omitempty" xml:"id"`

	// 订阅ID
	SubscriberId *string `json:"subscriberId,omitempty" xml:"subscriberId"`

	// 订阅url
	Url *string `json:"url,omitempty" xml:"url"`

	// 连接状态： 0表示已启用 ；1表示未启动； 2表示已锁定
	Status         *int32 `json:"status,omitempty" xml:"status"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowWebHookConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWebHookConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowWebHookConfigResponse", string(data)}, " ")
}
