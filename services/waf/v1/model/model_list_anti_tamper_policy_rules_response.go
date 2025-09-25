package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntiTamperPolicyRulesResponse Response Object
type ListAntiTamperPolicyRulesResponse struct {

	// 防泄漏规则数量
	Total *int32 `json:"total,omitempty"`

	// 防泄漏规则列表
	Items          *[]LeakageListInfo `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAntiTamperPolicyRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAntiTamperPolicyRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAntiTamperPolicyRulesResponse", string(data)}, " ")
}
