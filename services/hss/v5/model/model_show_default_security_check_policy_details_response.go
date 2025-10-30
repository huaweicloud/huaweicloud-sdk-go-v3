package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDefaultSecurityCheckPolicyDetailsResponse Response Object
type ShowDefaultSecurityCheckPolicyDetailsResponse struct {

	// **参数解释** 检查项列表总数 **取值范围** 取值0-2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 检查项列表
	CheckDetails   *[]SecurityCheckPolicyDetailInfoResponseInfo `json:"check_details,omitempty"`
	HttpStatusCode int                                          `json:"-"`
}

func (o ShowDefaultSecurityCheckPolicyDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultSecurityCheckPolicyDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowDefaultSecurityCheckPolicyDetailsResponse", string(data)}, " ")
}
