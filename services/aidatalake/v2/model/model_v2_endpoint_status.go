package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V2EndpointStatus **参数解释**：Endpoint状态。 **取值范围**：   - CREATING：创建中。   - RUNNING：运行中。   - CREATE_FAILED：创建失败。   - DELETING：删除中。   - DELETED：已删除。   - UPDATING：更新中。   - UPDATE_FAILED：更新失败。   - DELETE_FAILED：删除失败。   - FROZEN：冻结。   - TRIAL_ENDED：试用结束。   - ON_DEMAND_FROZEN：按需冻结。
type V2EndpointStatus struct {
}

func (o V2EndpointStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2EndpointStatus struct{}"
	}

	return strings.Join([]string{"V2EndpointStatus", string(data)}, " ")
}
