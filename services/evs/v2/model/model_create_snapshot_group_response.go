package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSnapshotGroupResponse Response Object
type CreateSnapshotGroupResponse struct {

	// **参数解释** 任务ID。  - 如果需要查询job的状态，请参考：[\"[查询job的状态](https://support.huaweicloud.com/api-evs/evs_04_0054.html)\"](tag:hws)[\"[查询job的状态](https://support.huaweicloud.com/intl/zh-cn/api-evs/evs_04_0054.html)\"](tag:hws_hk)。 **取值范围** 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释** 快照一致性组ID。 **取值范围** 不涉及。
	SnapshotGroupId *string `json:"snapshot_group_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o CreateSnapshotGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSnapshotGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateSnapshotGroupResponse", string(data)}, " ")
}
