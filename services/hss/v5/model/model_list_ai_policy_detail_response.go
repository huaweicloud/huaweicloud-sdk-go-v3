package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiPolicyDetailResponse Response Object
type ListAiPolicyDetailResponse struct {

	// **参数解释**: 策略详情 **取值范围**: 字符长度0-65535位
	Content        *string `json:"content,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAiPolicyDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiPolicyDetailResponse struct{}"
	}

	return strings.Join([]string{"ListAiPolicyDetailResponse", string(data)}, " ")
}
