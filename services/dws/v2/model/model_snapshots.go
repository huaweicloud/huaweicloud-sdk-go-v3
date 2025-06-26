package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Snapshots **参数解释**： 快照对象列表。 **取值范围**： 不涉及。
type Snapshots struct {

	// **参数解释**： 快照ID。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 快照名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 快照描述。 **取值范围**： 不涉及。
	Description string `json:"description"`

	// **参数解释**： 快照创建的日期时间，格式为 ISO8601: YYYY-MM-DDThh:mm:ssZ。 **取值范围**： 不涉及。
	Started string `json:"started"`

	// **参数解释**： 快照完成的日期时间，格式为 ISO8601: YYYY-MM-DDThh:mm:ssZ。 **取值范围**： 不涉及。
	Finished string `json:"finished"`

	// **参数解释**： 快照大小，单位 GB。 **取值范围**： 不涉及。
	Size float64 `json:"size"`

	// **参数解释**： 快照状态。 **取值范围**： - CREATING：创建中。 - AVAILABLE：可用。 - UNAVAILABLE：不可用。 - FROZEN： 普通冻结。 - POLICE_FROZEN： 公安冻结。
	Status string `json:"status"`

	// **参数解释**： 快照创建类型。 **取值范围**： 不涉及。
	Type string `json:"type"`

	// **参数解释**： 快照对应的集群ID。 **取值范围**： 不涉及。
	ClusterId string `json:"cluster_id"`
}

func (o Snapshots) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Snapshots struct{}"
	}

	return strings.Join([]string{"Snapshots", string(data)}, " ")
}
