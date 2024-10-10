package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWhiteBlackIpRuleResponse Response Object
type ListWhiteBlackIpRuleResponse struct {

	// total
	Total *int32 `json:"total,omitempty"`

	// 名单列表
	Ips            *[]BwListIps `json:"ips,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListWhiteBlackIpRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWhiteBlackIpRuleResponse struct{}"
	}

	return strings.Join([]string{"ListWhiteBlackIpRuleResponse", string(data)}, " ")
}
