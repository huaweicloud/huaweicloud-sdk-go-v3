package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlarmConfigResponse Response Object
type DeleteAlarmConfigResponse struct {

	// SMN的topic urn
	TopicUrn       *string `json:"topic_urn,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAlarmConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteAlarmConfigResponse", string(data)}, " ")
}
