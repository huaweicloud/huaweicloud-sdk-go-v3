package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmConfigResponse Response Object
type ShowAlarmConfigResponse struct {

	// SMN的topic urn
	TopicUrn       *string `json:"topic_urn,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAlarmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmConfigResponse", string(data)}, " ")
}
