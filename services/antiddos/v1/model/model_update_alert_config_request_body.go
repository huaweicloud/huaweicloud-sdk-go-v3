package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAlertConfigRequestBody struct {
	// 告警群组描述。

	DisplayName string `json:"display_name"`
	// 告警群组的唯一标识。

	TopicUrn string `json:"topic_urn"`

	WarnConfig *UpdateAlertConfigRequestBodyWarnConfig `json:"warn_config"`
}

func (o UpdateAlertConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlertConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAlertConfigRequestBody", string(data)}, " ")
}
