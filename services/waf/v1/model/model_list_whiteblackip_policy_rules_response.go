package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWhiteblackipPolicyRulesResponse Response Object
type ListWhiteblackipPolicyRulesResponse struct {

	// 黑白名单规则条数
	Total *int32 `json:"total,omitempty"`

	// 黑白名单规则列表信息
	Items *[]WhiteBlackIpResponseBody `json:"items,omitempty"`

	// ip地址总数
	Size           *int32 `json:"size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListWhiteblackipPolicyRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWhiteblackipPolicyRulesResponse struct{}"
	}

	return strings.Join([]string{"ListWhiteblackipPolicyRulesResponse", string(data)}, " ")
}
