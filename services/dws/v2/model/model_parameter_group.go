package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParameterGroup **参数解释**： 集群所关联的参数组信息。 **取值范围**： 不涉及。
type ParameterGroup struct {

	// **参数解释**： 参数组ID。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 参数组名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 集群参数状态。 **取值范围**： - In-Sync：已同步 - Applying：应用中 - Pending-Reboot：需重启生效 - Sync-Failure：应用失败
	Status string `json:"status"`
}

func (o ParameterGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParameterGroup struct{}"
	}

	return strings.Join([]string{"ParameterGroup", string(data)}, " ")
}
