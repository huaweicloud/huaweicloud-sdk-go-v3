package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupsStatisticsResponse Response Object
type ShowBackupsStatisticsResponse struct {

	// **参数解释**： 已备份服务器数量 **取值范围**： 最小值0，最大值2147483647
	CountBackedupServer *int32 `json:"count_backedup_server,omitempty"`

	// **参数解释**： 已恢复的备份任务数量 **取值范围**： 最小值0，最大值2147483647
	CountRestoredTask *int32 `json:"count_restored_task,omitempty"`

	// **参数解释**： 已使用容量 **取值范围**： 最小值0，最大值2147483647
	CountUsedCapacity *int32 `json:"count_used_capacity,omitempty"`

	// **参数解释**： 总容量 **取值范围**： 最小值0，最大值2147483647
	CountTotalCapacity *int32 `json:"count_total_capacity,omitempty"`

	// **参数解释**： 存储库状态 **取值范围**： 字符长度0-64位
	VaultStatus    *string `json:"vault_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBackupsStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupsStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupsStatisticsResponse", string(data)}, " ")
}
