package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAiOpsRequest Request Object
type CreateAiOpsRequest struct {

	// **参数解释**： 指定操作的集群ID。获取方法请参见[获取集群ID](css_03_0101.xml)。 **约束限制**： 不涉及 **取值范围**： 集群ID。 **默认取值**： 不涉及
	ClusterId string `json:"cluster_id"`

	Body *CreateAiOpsRequestBody `json:"body,omitempty"`
}

func (o CreateAiOpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAiOpsRequest struct{}"
	}

	return strings.Join([]string{"CreateAiOpsRequest", string(data)}, " ")
}
