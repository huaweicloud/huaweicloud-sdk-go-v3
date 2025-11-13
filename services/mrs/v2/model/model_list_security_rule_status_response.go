package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityRuleStatusResponse Response Object
type ListSecurityRuleStatusResponse struct {

	// 获取通信安全授权状态详情
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSecurityRuleStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityRuleStatusResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityRuleStatusResponse", string(data)}, " ")
}
