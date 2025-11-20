package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDefaultSecurityCheckPolicyResponse Response Object
type ShowDefaultSecurityCheckPolicyResponse struct {

	// **参数解释**: 策略详情 **取值范围**: 最小值0，最大值10241
	Content        *string `json:"content,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDefaultSecurityCheckPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultSecurityCheckPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowDefaultSecurityCheckPolicyResponse", string(data)}, " ")
}
