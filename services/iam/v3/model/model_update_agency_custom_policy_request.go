package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateAgencyCustomPolicyRequest struct {
	// 待修改的自定义策略ID，获取方式请参见：[自定义策略ID](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=IAM&api=ListCustomPolicies)。

	RoleId string `json:"role_id"`

	Body *UpdateAgencyCustomPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateAgencyCustomPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAgencyCustomPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateAgencyCustomPolicyRequest", string(data)}, " ")
}
