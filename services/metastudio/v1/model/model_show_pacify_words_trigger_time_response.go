package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPacifyWordsTriggerTimeResponse Response Object
type ShowPacifyWordsTriggerTimeResponse struct {

	// 安抚话术等待触发时长，单位毫秒
	TriggerTime *int32 `json:"trigger_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPacifyWordsTriggerTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPacifyWordsTriggerTimeResponse struct{}"
	}

	return strings.Join([]string{"ShowPacifyWordsTriggerTimeResponse", string(data)}, " ")
}
