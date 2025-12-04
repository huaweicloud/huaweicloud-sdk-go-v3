package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetConfigurationResponse Response Object
type ResetConfigurationResponse struct {

	// **参数解释**：  参数组模板名称。  **约束限制**：  不涉及。
	ConfigName *string `json:"config_name,omitempty"`

	// **参数解释**：  是否涉及重启。（当前该字段不使用，默认为false）  **约束限制**：  不涉及。
	NeedRestart    *bool `json:"need_restart,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ResetConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ResetConfigurationResponse", string(data)}, " ")
}
