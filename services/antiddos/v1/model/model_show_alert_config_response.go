package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAlertConfigResponse struct {

	// 告警群组的唯一标识
	TopicUrn *string `json:"topic_urn,omitempty" xml:"topic_urn"`

	// 告警群组描述
	DisplayName *string `json:"display_name,omitempty" xml:"display_name"`

	WarnConfig     *AlertConfigRespWarnConfig `json:"warn_config,omitempty" xml:"warn_config"`
	HttpStatusCode int                        `json:"-"`
}

func (o ShowAlertConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlertConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAlertConfigResponse", string(data)}, " ")
}
