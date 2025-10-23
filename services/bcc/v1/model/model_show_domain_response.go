package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainResponse Response Object
type ShowDomainResponse struct {

	// 是否完成启用
	Enabled *bool `json:"enabled,omitempty"`

	// 数据是否同步完成
	DataSyncFinished *bool `json:"data_sync_finished,omitempty"`

	// 数据同步进度百分比，只有data_sync_finished为false时才有值
	DataSyncPercent *int32 `json:"data_sync_percent,omitempty"`

	// 是否创建资源等级
	ResourceLevelCreated *bool `json:"resource_level_created,omitempty"`

	// 是否创建合规规则
	BackupComplianceRuleCreated *bool `json:"backup_compliance_rule_created,omitempty"`

	// 是否绑定合规规则
	BackupComplianceRuleBound *bool `json:"backup_compliance_rule_bound,omitempty"`
	HttpStatusCode            int   `json:"-"`
}

func (o ShowDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainResponse", string(data)}, " ")
}
