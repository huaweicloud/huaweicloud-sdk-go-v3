package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOperationalTaskConfigRequest Request Object
type UpdateOperationalTaskConfigRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	Body *OperationalTaskConfiguration `json:"body,omitempty"`
}

func (o UpdateOperationalTaskConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOperationalTaskConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateOperationalTaskConfigRequest", string(data)}, " ")
}
