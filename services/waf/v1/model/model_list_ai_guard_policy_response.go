package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiGuardPolicyResponse Response Object
type ListAiGuardPolicyResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 策略列表
	Items          *[]AiGuardPolicyItem `json:"items,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListAiGuardPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiGuardPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListAiGuardPolicyResponse", string(data)}, " ")
}
