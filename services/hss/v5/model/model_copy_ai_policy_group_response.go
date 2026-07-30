package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyAiPolicyGroupResponse Response Object
type CopyAiPolicyGroupResponse struct {

	// **参数解释**： 策略组ID **取值范围**： 字符长度1-20位
	GroupId        *string `json:"group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyAiPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyAiPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"CopyAiPolicyGroupResponse", string(data)}, " ")
}
