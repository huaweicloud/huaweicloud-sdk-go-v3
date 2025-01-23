package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBackupOffsitePolicyInfo 备份策略信息。
type UpdateBackupOffsitePolicyInfo struct {

	// 是否开启跨区域全量备份。 - true：开启跨区域全量备份。 - false：关闭跨区域全量备份。
	OpenAutoBackup bool `json:"open_auto_backup"`

	// 是否开启跨区域增量备份。 - true：开启跨区域增量备份，当open_auto_backup开启时才可以开启。 - false：关闭跨区域增量备份。
	OpenIncrementalBackup bool `json:"open_incremental_backup"`

	// 设置跨区域备份策略的目标project ID。    租户在某一Region下的project ID。                              获取方法请参见[获取项目ID](https://support.huaweicloud.com/api-taurusdb/taurusdb_10_0004.html)。
	DestinationProjectId string `json:"destination_project_id"`

	// 设置跨区域备份策略的目标区域。
	DestinationRegion string `json:"destination_region"`

	// 指定已生成的备份文件可以保存的天数。  取值范围：1～1825。
	KeepDays int32 `json:"keep_days"`
}

func (o UpdateBackupOffsitePolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBackupOffsitePolicyInfo struct{}"
	}

	return strings.Join([]string{"UpdateBackupOffsitePolicyInfo", string(data)}, " ")
}
