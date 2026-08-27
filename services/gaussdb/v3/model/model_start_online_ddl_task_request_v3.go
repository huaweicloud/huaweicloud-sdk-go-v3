package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartOnlineDdlTaskRequestV3 **参数解释**：  开启无锁变更任务请求体。
type StartOnlineDdlTaskRequestV3 struct {

	// **参数解释**：  是否开启自动清理临时表。  **约束限制**：  不涉及。  **取值范围**： - true：开启自动清理临时表。 - false：关闭自动清理临时表。  **默认取值**：  false。
	AutoClear *bool `json:"auto_clear,omitempty"`

	// **参数解释**：  无锁变更任务详细内容。  **约束限制**：  不涉及。
	TaskContent []StartOnlineTaskContentItem `json:"task_content"`
}

func (o StartOnlineDdlTaskRequestV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartOnlineDdlTaskRequestV3 struct{}"
	}

	return strings.Join([]string{"StartOnlineDdlTaskRequestV3", string(data)}, " ")
}
