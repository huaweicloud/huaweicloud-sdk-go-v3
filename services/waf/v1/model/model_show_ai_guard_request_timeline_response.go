package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiGuardRequestTimelineResponse Response Object
type ShowAiGuardRequestTimelineResponse struct {

	// AI Guard请求时间线响应
	Body *[]AiGuardTimeLineItem `json:"body,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAiGuardRequestTimelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiGuardRequestTimelineResponse struct{}"
	}

	return strings.Join([]string{"ShowAiGuardRequestTimelineResponse", string(data)}, " ")
}
