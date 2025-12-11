package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchGaussMySqlConfigurationResponse Response Object
type SwitchGaussMySqlConfigurationResponse struct {

	// **参数解释**：  应用参数模板的任务ID。  **取值范围**：  不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：  参数模板的名称。  **取值范围**：  支持Default-TaurusDB V2.0和用户自定义参数模板，其中Default-TaurusDB V2.0表示TaurusDB系统默认参数模板。
	ParamGroupName *string `json:"param_group_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchGaussMySqlConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlConfigurationResponse struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlConfigurationResponse", string(data)}, " ")
}
