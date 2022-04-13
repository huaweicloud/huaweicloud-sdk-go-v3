package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SwitchAgentResponse struct {
	// 响应状态

	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAgentResponse struct{}"
	}

	return strings.Join([]string{"SwitchAgentResponse", string(data)}, " ")
}
