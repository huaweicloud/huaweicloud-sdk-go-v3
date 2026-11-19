package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityCheckPolicyGroupHostsResponse Response Object
type ListSecurityCheckPolicyGroupHostsResponse struct {

	// **参数解释**: 基线策略关联主机总数量 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： 主机信息列表 **取值范围**： 不涉及
	DataList       *[]SecurityCheckPolicyGroupHostInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListSecurityCheckPolicyGroupHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityCheckPolicyGroupHostsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityCheckPolicyGroupHostsResponse", string(data)}, " ")
}
