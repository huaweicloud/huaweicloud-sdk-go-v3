package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchAgentNewResponse Response Object
type SwitchAgentNewResponse struct {

	// 操作结果  - SUCCESS: 成功
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchAgentNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchAgentNewResponse struct{}"
	}

	return strings.Join([]string{"SwitchAgentNewResponse", string(data)}, " ")
}
