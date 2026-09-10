package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ArtifactsRsp struct {

	// **参数解释**： 创建时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**： 标签列表。 **约束限制**： 产物列表不能超过10条。 **取值范围**： 不涉及 **默认取值**： 不涉及
	Filenames *[]string `json:"filenames,omitempty"`

	// **参数解释**： 绑定状态。 **约束限制**： 不涉及 **取值范围**： * requirement_analyzer：构建需求文档。 * modeling：构建数学模型。 * data：校验模型数据。 * solver：求解数学模型。 * report：业务辅助分析。 * business_planner：构建需求文档 * data_agent：原始数据处理 * vrp：路径规划求解 **默认取值**： 不涉及
	StageName *string `json:"stage_name,omitempty"`
}

func (o ArtifactsRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArtifactsRsp struct{}"
	}

	return strings.Join([]string{"ArtifactsRsp", string(data)}, " ")
}
