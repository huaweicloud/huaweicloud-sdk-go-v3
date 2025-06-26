package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterConfiguration **参数解释**： 集群所关联的参数组详情。 **取值范围**： 不涉及。
type ClusterConfiguration struct {

	// **参数解释**： 参数组ID。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 参数组名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 参数组类型。 **取值范围**： 不涉及。
	Type string `json:"type"`

	// **参数解释**： 集群参数状态。 **取值范围**： In-Sync：已同步。 Applying：应用中。 Pending-Reboot：需重启生效。 Sync-Failure：应用失败。
	Status string `json:"status"`

	// **参数解释**： 参数应用失败原因。 **取值范围**： 不涉及。
	FailReason string `json:"fail_reason"`
}

func (o ClusterConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterConfiguration struct{}"
	}

	return strings.Join([]string{"ClusterConfiguration", string(data)}, " ")
}
