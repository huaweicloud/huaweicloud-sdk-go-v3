package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartInfo 启动任务请求体。
type StartInfo struct {

	// **参数解释**：  迁移任务ID，用于唯一标识一个数据复制任务，在查询、修改、删除等操作中作为任务的识别凭证。  **约束限制**：  不涉及。  **取值范围**：  UUID格式字符串。  **默认取值**：  不涉及
	JobId string `json:"job_id"`

	// **参数解释**：  是否只初始化任务，指定启动任务时是否仅执行初始化而不启动数据迁移。  **约束限制**：  - 仅支持白名单用户使用，需要提交工单申请才能使用。  **取值范围**：  - true：只初始化任务 - false：正常启动任务  **默认取值**：  false
	IsOnlyInitTask *bool `json:"is_only_init_task,omitempty"`

	// **参数解释**：  是否在任务结束时自动创建对比任务，指定任务完成后是否自动触发数据对比以验证迁移一致性。  **约束限制**：  不涉及。  **取值范围**：  - true：自动创建对比任务 - false：不自动创建对比任务  **默认取值**：  true
	IsAutoCreateCompare *bool `json:"is_auto_create_compare,omitempty"`

	// **参数解释**：  任务启动时间，指定任务开始执行的时间，时间戳格式精确到毫秒。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及，取值为空代表立即启动
	StartTime *string `json:"start_time,omitempty"`
}

func (o StartInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInfo struct{}"
	}

	return strings.Join([]string{"StartInfo", string(data)}, " ")
}
