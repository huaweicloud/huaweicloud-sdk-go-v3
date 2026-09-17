package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutopilotReleaseHistoryRequest Request Object
type ShowAutopilotReleaseHistoryRequest struct {

	// **参数解释：** 模板实例名称。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name string `json:"name"`

	// **参数解释：** 模板实例所在的命名空间。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Namespace string `json:"namespace"`

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`
}

func (o ShowAutopilotReleaseHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutopilotReleaseHistoryRequest struct{}"
	}

	return strings.Join([]string{"ShowAutopilotReleaseHistoryRequest", string(data)}, " ")
}
