package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupPolicyInfo **参数解释**：  备份策略信息。
type BackupPolicyInfo struct {

	// **参数解释**：  指定已生成的备份文件可以保存的天数。  **取值范围**：  1-732。 您也可以联系客服申请开通最大保留天数为3660。
	RetentionDays int32 `json:"retention_days"`

	// **参数解释**：  备份周期配置。  **取值范围**：  格式必须为“日期 月份 星期”形式的Cron表达式，时区为UTC时区。 日期支持填写1~31、特殊字符*（表示任意值）、特殊字符L（表示最后一天）。填写1~31或L时支持填写多个，需以逗号隔开。 月份支持填写1~12、特殊字符*（表示任意值）。 星期支持填写1~7，需以逗号隔开。
	Period string `json:"period"`

	// **参数解释**:  自动备份策略类型。  **取值范围**：   - base：表示基础策略。   - sparse：表示稀疏策略。
	PolicyType string `json:"policy_type"`
}

func (o BackupPolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupPolicyInfo struct{}"
	}

	return strings.Join([]string{"BackupPolicyInfo", string(data)}, " ")
}
