package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIncidentTaskResponse Response Object
type ShowIncidentTaskResponse struct {

	// 服务标识
	ProviderCode string `json:"provider_code"`

	// 请求响应代码，范围：0000~9999，正常时取值：0
	ErrorCode string `json:"error_code"`

	// 请求响应描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 任务列表
	Data           *[]CocTaskDetailV2 `json:"data,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowIncidentTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIncidentTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowIncidentTaskResponse", string(data)}, " ")
}
