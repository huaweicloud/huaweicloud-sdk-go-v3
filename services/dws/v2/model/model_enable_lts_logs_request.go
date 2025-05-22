package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableLtsLogsRequest Request Object
type EnableLtsLogsRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`
}

func (o EnableLtsLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableLtsLogsRequest struct{}"
	}

	return strings.Join([]string{"EnableLtsLogsRequest", string(data)}, " ")
}
