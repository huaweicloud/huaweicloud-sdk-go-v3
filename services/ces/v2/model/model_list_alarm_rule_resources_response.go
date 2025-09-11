package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmRuleResourcesResponse Response Object
type ListAlarmRuleResourcesResponse struct {

	// **参数解释**： 资源信息。
	Resources *[][]DimensionResp `json:"resources,omitempty"`

	// **参数解释**： 资源总数。 **取值范围**： [0,2147483647]
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAlarmRuleResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmRuleResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListAlarmRuleResourcesResponse", string(data)}, " ")
}
