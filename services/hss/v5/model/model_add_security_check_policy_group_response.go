package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSecurityCheckPolicyGroupResponse Response Object
type AddSecurityCheckPolicyGroupResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddSecurityCheckPolicyGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSecurityCheckPolicyGroupResponse struct{}"
	}

	return strings.Join([]string{"AddSecurityCheckPolicyGroupResponse", string(data)}, " ")
}
