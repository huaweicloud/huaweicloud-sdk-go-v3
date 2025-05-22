package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOfflineKeyAnalysisRequestBody 创建离线全量key分析请求体
type CreateOfflineKeyAnalysisRequestBody struct {

	// **参数解释**： 实例节点ID，可通过DCS控制台进入实例的节点管理界面查看。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**： 实例备份ID，可通过DCS控制台进入实例的备份与恢复界面查看。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BackupId *string `json:"backup_id,omitempty"`
}

func (o CreateOfflineKeyAnalysisRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOfflineKeyAnalysisRequestBody struct{}"
	}

	return strings.Join([]string{"CreateOfflineKeyAnalysisRequestBody", string(data)}, " ")
}
