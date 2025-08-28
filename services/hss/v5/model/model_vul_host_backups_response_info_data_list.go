package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulHostBackupsResponseInfoDataList struct {

	// **参数解释**: 备份时间（时间戳，单位为毫秒） **取值范围**: 取值0-9223372036854775807
	BackupTime *int64 `json:"backup_time,omitempty"`

	// **参数解释**: 备份id **取值范围**: 字符长度1-128位
	BackupId *string `json:"backup_id,omitempty"`

	// **参数解释**: 备份名称 **取值范围**: 字符长度1-64位
	BackupName *string `json:"backup_name,omitempty"`

	// **参数解释**: 备份状态 **取值范围**: - available：可用 - protecting：备份中 - deleting：删除中 - restoring：恢复中 - error：错误 - waiting_protect：待备份 - waiting_delete：待删除 - waiting_restore：待恢复
	Status *string `json:"status,omitempty"`
}

func (o VulHostBackupsResponseInfoDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulHostBackupsResponseInfoDataList struct{}"
	}

	return strings.Join([]string{"VulHostBackupsResponseInfoDataList", string(data)}, " ")
}
