package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceTaskResponse Response Object
type ShowInstanceTaskResponse struct {

	// 任务Id
	TaskId *string `json:"task_id,omitempty"`

	// **参数说明**：实例任务类型。 **取值范围**： - CREATE：创建实例任务 - MODIFY：实例规格变更任务 - DELETE：实例删除任务 - FREEZE：实例冻结任务 - UNFREEZE：实例解冻任务 - UPDATE_ACCESS_CONFIG：修改实例接入信息任务 - UPDATE_ALLOW_LISTS： 修改实例接入白名单任务 - OPEN_SNAT：启用实例SNAT配置任务
	Type *string `json:"type,omitempty"`

	// **参数说明**：任务状态。 **取值范围**： - PENDING：等待执行 - RUNNING：执行中 - SUCCESS：执行成功 - FAILED：执行失败
	Status *string `json:"status,omitempty"`

	// **参数说明**：任务状态描述。
	StatusDetail *string `json:"status_detail,omitempty"`

	// **参数说明**：实例任务的创建时间。格式为：\"yyyy-MM-dd'T'HH:mm:ss.SSS'Z'\"
	CreateTime *string `json:"create_time,omitempty"`

	// **参数说明**：实例任务的开始时间。格式为：\"yyyy-MM-dd'T'HH:mm:ss.SSS'Z'\"
	StartTime *string `json:"start_time,omitempty"`

	// **参数说明**：实例任务的结束时间。格式为：\"yyyy-MM-dd'T'HH:mm:ss.SSS'Z'\"
	EndTime *string `json:"end_time,omitempty"`

	TargetConfig *TargetConfig `json:"target_config,omitempty"`

	OperateWindow *OperateWindow `json:"operate_window,omitempty"`

	// **参数说明**：任务进度\"
	Progress       *int32 `json:"progress,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceTaskResponse", string(data)}, " ")
}
