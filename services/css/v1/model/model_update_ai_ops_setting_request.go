package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiOpsSettingRequest Request Object
type UpdateAiOpsSettingRequest struct {

	// **参数解释**： 指定查询的集群ID。获取方法请参见[获取集群ID](css_03_0101.xml)。 **约束限制**： 不涉及 **取值范围**： 集群ID。 **默认取值**： 不涉及
	ClusterId string `json:"cluster_id"`

	Body *AiOpsSetting `json:"body,omitempty"`
}

func (o UpdateAiOpsSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiOpsSettingRequest struct{}"
	}

	return strings.Join([]string{"UpdateAiOpsSettingRequest", string(data)}, " ")
}
