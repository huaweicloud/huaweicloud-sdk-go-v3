package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingJobRoutePlanResponse Response Object
type ShowTrainingJobRoutePlanResponse struct {

	// **参数解释**：训练作业ID。 **取值范围**：不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：路由规划状态。 **取值范围**： - success：路由规划成功 - failed：路由规划未执行或不满足条件，返回默认rank映射
	Status *string `json:"status,omitempty"`

	// **参数解释**：rank映射结果，格式为\"newRankId-workerId\"，多个映射项之间以英文逗号分隔。 **约束限制**：当status为failed时，返回基于作业规格计算的默认顺序映射。 **取值范围**：不涉及。
	RankMapping    *string `json:"rank_mapping,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTrainingJobRoutePlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingJobRoutePlanResponse struct{}"
	}

	return strings.Join([]string{"ShowTrainingJobRoutePlanResponse", string(data)}, " ")
}
