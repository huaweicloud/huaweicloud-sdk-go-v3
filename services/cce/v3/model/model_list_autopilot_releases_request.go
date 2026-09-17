package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutopilotReleasesRequest Request Object
type ListAutopilotReleasesRequest struct {

	// 集群ID，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ClusterId string `json:"cluster_id"`

	// **参数解释：** 模板ID。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ChartId *string `json:"chart_id,omitempty"`

	// **参数解释：** 模板对应的命名空间。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Namespace *string `json:"namespace,omitempty"`
}

func (o ListAutopilotReleasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutopilotReleasesRequest struct{}"
	}

	return strings.Join([]string{"ListAutopilotReleasesRequest", string(data)}, " ")
}
