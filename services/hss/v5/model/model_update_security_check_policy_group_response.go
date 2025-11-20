package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityCheckPolicyGroupResponse Response Object
type UpdateSecurityCheckPolicyGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSecurityCheckPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityCheckPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecurityCheckPolicyGroupResponse", string(data)}, " ")
}
