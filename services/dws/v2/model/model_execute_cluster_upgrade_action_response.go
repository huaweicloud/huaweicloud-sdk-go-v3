package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteClusterUpgradeActionResponse Response Object
type ExecuteClusterUpgradeActionResponse struct {

	// **参数解释**： 集群升级任务ID。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteClusterUpgradeActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteClusterUpgradeActionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteClusterUpgradeActionResponse", string(data)}, " ")
}
