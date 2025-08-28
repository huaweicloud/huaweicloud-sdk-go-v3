package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVulBackupStatisticsResponse Response Object
type ShowVulBackupStatisticsResponse struct {

	// **参数解释**: 本次漏洞处理的备份信息id **取值范围**: 字符长度1-128位
	BackupInfoId *string `json:"backup_info_id,omitempty"`

	// **参数解释**: 本次漏洞处理可进行备份的主机数量 **取值范围**: 字符长度0-1000000位
	BackupHostNum *int32 `json:"backup_host_num,omitempty"`

	// **参数解释**: 本次漏洞处理不可进行备份的主机数量 **取值范围**: 字符长度0-1000000位
	UnableBackupHostNum *int32 `json:"unable_backup_host_num,omitempty"`
	HttpStatusCode      int    `json:"-"`
}

func (o ShowVulBackupStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVulBackupStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowVulBackupStatisticsResponse", string(data)}, " ")
}
