package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteAiPolicyGroupsRequestInfo struct {

	// **参数解释**： 策略组ID列表 **约束限制**： 必填 **取值范围**： 1-200个策略组ID **默认取值**： 不涉及
	GroupIdList []string `json:"group_id_list"`
}

func (o DeleteAiPolicyGroupsRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAiPolicyGroupsRequestInfo struct{}"
	}

	return strings.Join([]string{"DeleteAiPolicyGroupsRequestInfo", string(data)}, " ")
}
