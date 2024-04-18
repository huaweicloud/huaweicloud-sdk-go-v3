package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowErrorLogSwitchStatusRequest Request Object
type ShowErrorLogSwitchStatusRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowErrorLogSwitchStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowErrorLogSwitchStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowErrorLogSwitchStatusRequest", string(data)}, " ")
}
