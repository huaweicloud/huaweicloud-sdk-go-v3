package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectionPolicyResponse Response Object
type ListProtectionPolicyResponse struct {

	// **参数解释**: 策略总数 **取值范围**: 取值0-2097152
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 查询防护策略列表 **取值范围**: 取值0-2097152
	DataList       *[]ProtectionPolicyInfo `json:"data_list,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListProtectionPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectionPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListProtectionPolicyResponse", string(data)}, " ")
}
