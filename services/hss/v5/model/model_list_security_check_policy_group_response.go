package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityCheckPolicyGroupResponse Response Object
type ListSecurityCheckPolicyGroupResponse struct {

	// **参数解释** 配置检测总数 **取值范围** 取值0-1000000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 数据列表
	DataList       *[]SecurityCheckPolicyGroupInfoResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o ListSecurityCheckPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityCheckPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityCheckPolicyGroupResponse", string(data)}, " ")
}
