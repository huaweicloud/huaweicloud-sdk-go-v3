package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupTriggerRequestInfo1 **参数解释**: 策略时间调度规则。 **约束限制**: 不涉及 **取值范围**: 不涉及  **默认取值**: 不涉及
type BackupTriggerRequestInfo1 struct {
	Properties *BackupTriggerPropertiesRequestInfo1 `json:"properties,omitempty"`
}

func (o BackupTriggerRequestInfo1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupTriggerRequestInfo1 struct{}"
	}

	return strings.Join([]string{"BackupTriggerRequestInfo1", string(data)}, " ")
}
