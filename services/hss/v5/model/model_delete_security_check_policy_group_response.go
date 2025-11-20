package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityCheckPolicyGroupResponse Response Object
type DeleteSecurityCheckPolicyGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityCheckPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityCheckPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityCheckPolicyGroupResponse", string(data)}, " ")
}
