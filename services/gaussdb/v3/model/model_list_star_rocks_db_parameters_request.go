package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStarRocksDbParametersRequest Request Object
type ListStarRocksDbParametersRequest struct {

	// StarRocks实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage string `json:"X-Language"`

	// **参数解释**：  新增子任务的场景，用于区分库参数是否支持修改。  **约束限制**：  非必填。  **取值范围**：  不涉及。  **默认值**：  不涉及。
	AddTaskScenario *string `json:"add_task_scenario,omitempty"`

	// **参数解释**：  新增子任务相应的主任务名。  **约束限制**：  非必填。  **取值范围**：  不涉及。  **默认值**：  不涉及。
	MainTaskName *string `json:"main_task_name,omitempty"`
}

func (o ListStarRocksDbParametersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStarRocksDbParametersRequest struct{}"
	}

	return strings.Join([]string{"ListStarRocksDbParametersRequest", string(data)}, " ")
}
