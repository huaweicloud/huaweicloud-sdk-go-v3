package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCheckRulesInfoResponse Response Object
type ListCheckRulesInfoResponse struct {

	// **参数解释**： 总数 **取值范围**： 最小值0，最大值2147483547
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 镜像的配置检查信息 **取值范围**: 最小值0，最大值10241
	DataList       *[]ImageCheckRuleResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListCheckRulesInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCheckRulesInfoResponse struct{}"
	}

	return strings.Join([]string{"ListCheckRulesInfoResponse", string(data)}, " ")
}
