package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreDuplicationRequestInfo struct {

	// **参数解释**: 恢复的目标虚拟机ID **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	ServerId *string `json:"server_id,omitempty"`

	// **参数解释**： 恢复后是否开机, 默认开机。 **约束限制**: 不涉及 **取值范围**: - true：开机 - false: 不开机 **默认取值**: true
	PowerOn *bool `json:"power_on,omitempty"`

	// **参数解释**： 恢复的映射关系（整机恢复时必填，卷恢复时可选但是不会用到填写的值） **取值范围**: 不涉及
	Mappings *[]BackupRestoreServerMappingInfo `json:"mappings,omitempty"`
}

func (o RestoreDuplicationRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreDuplicationRequestInfo struct{}"
	}

	return strings.Join([]string{"RestoreDuplicationRequestInfo", string(data)}, " ")
}
