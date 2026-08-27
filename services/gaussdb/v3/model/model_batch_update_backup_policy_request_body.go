package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateBackupPolicyRequestBody 批量设置同区域备份策略请求体
type BatchUpdateBackupPolicyRequestBody struct {

	// **参数解释**：  需要设置备份策略的实例ID列表。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  列表数量小于等于50。
	InstanceIds []string `json:"instance_ids"`

	BackupPolicy *MysqlBackupPolicyInfo `json:"backup_policy"`
}

func (o BatchUpdateBackupPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateBackupPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateBackupPolicyRequestBody", string(data)}, " ")
}
