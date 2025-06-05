package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupTriggerInfo 备份:策略时间调度规则
type BackupTriggerInfo struct {

	// **参数解释**: 调度器id **取值范围**: 字符长度0-256
	Id *string `json:"id,omitempty"`

	// **参数解释**: 调度器名称 **取值范围**: 字符长度0-256
	Name *string `json:"name,omitempty"`

	// **参数解释**: 调度器类型,目前只支持 time,定时调度。 **取值范围**: 字符长度0-256
	Type *string `json:"type,omitempty"`

	Properties *BackupTriggerPropertiesInfo `json:"properties,omitempty"`
}

func (o BackupTriggerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupTriggerInfo struct{}"
	}

	return strings.Join([]string{"BackupTriggerInfo", string(data)}, " ")
}
