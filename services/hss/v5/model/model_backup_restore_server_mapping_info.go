package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackupRestoreServerMappingInfo struct {

	// **参数解释**: 卷备份ID，可以通过控制台或者“查询指定备份”接口获取。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	BackupId *string `json:"backup_id,omitempty"`

	// **参数解释**: 待恢复目标卷ID **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	VolumeId *string `json:"volume_id,omitempty"`
}

func (o BackupRestoreServerMappingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupRestoreServerMappingInfo struct{}"
	}

	return strings.Join([]string{"BackupRestoreServerMappingInfo", string(data)}, " ")
}
