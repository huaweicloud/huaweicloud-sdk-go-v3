package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResumePauseAiPolicyGroupsRequestInfo struct {

	// **参数解释**: 是否启用 **约束限制**: 必填 **取值范围**: - false：否 - true：是  **默认取值**: 不涉及
	Enabled bool `json:"enabled"`

	// **参数解释**： 策略组ID列表 **约束限制**： 必填 **取值范围**： 1-50个策略组ID **默认取值**： 不涉及
	GroupIdList []string `json:"group_id_list"`
}

func (o ResumePauseAiPolicyGroupsRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumePauseAiPolicyGroupsRequestInfo struct{}"
	}

	return strings.Join([]string{"ResumePauseAiPolicyGroupsRequestInfo", string(data)}, " ")
}
