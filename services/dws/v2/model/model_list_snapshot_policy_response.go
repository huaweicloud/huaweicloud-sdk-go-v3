package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotPolicyResponse Response Object
type ListSnapshotPolicyResponse struct {

	// **参数解释**： 保留天数。 **取值范围**： 大于等于0。
	KeepDay *int32 `json:"keep_day,omitempty"`

	// **参数解释**： 备份策略列表。 **取值范围**： 不涉及。
	BackupStrategies *[]BackupStrategyDetail `json:"backup_strategies,omitempty"`

	// **参数解释**： 备份设备，一般为OBS。 **取值范围**： 不涉及。
	DeviceName *string `json:"device_name,omitempty"`

	// **参数解释**： 服务IP。 **取值范围**： 不涉及。
	ServerIps *[]string `json:"server_ips,omitempty"`

	// **参数解释**： 服务端口。 **取值范围**： 不涉及。
	ServerPort *string `json:"server_port,omitempty"`

	// **参数解释**： 备份参数。 **取值范围**： 不涉及。
	BackupParam *string `json:"backup_param,omitempty"`

	// **参数解释**： 自动备份开关状态。 **取值范围**： true：已开启自动备份选项； false：已关闭自动备份选项；
	AutoBackup *bool `json:"auto_backup,omitempty"`

	// **参数解释**： 此策略下集群级快照最大数量。 **取值范围**： 大于等于0。
	BackupStrategyClusterTypeLimitNum *int32 `json:"backup_strategy_cluster_type_limit_num,omitempty"`

	// **参数解释**： 此策略下schema级快照最大数量。 **取值范围**： 大于等于0。
	BackupStrategySchemaTypeLimitNum *int32 `json:"backup_strategy_schema_type_limit_num,omitempty"`
	HttpStatusCode                   int    `json:"-"`
}

func (o ListSnapshotPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListSnapshotPolicyResponse", string(data)}, " ")
}
