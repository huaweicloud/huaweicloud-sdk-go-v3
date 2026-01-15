package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAiOpsRequest Request Object
type DeleteAiOpsRequest struct {

	// **参数解释**： 指定删除的集群ID。获取方法请参见[获取集群ID](css_03_0101.xml)。 **约束限制**： 不涉及 **取值范围**： 集群ID。 **默认取值**： 不涉及
	ClusterId string `json:"cluster_id"`

	// 指定检测任务ID。
	AiopsId string `json:"aiops_id"`
}

func (o DeleteAiOpsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAiOpsRequest struct{}"
	}

	return strings.Join([]string{"DeleteAiOpsRequest", string(data)}, " ")
}
