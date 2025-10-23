package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HandleIncidentResponse Response Object
type HandleIncidentResponse struct {

	// 服务标识
	ProviderCode *string `json:"provider_code,omitempty"`

	// 请求响应代码，范围：0000~9999，正常时取值：0
	ErrorCode *string `json:"error_code,omitempty"`

	// 请求响应描述
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o HandleIncidentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandleIncidentResponse struct{}"
	}

	return strings.Join([]string{"HandleIncidentResponse", string(data)}, " ")
}
