package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncSecurityGroupPoliciesResponse Response Object
type SyncSecurityGroupPoliciesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SyncSecurityGroupPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncSecurityGroupPoliciesResponse struct{}"
	}

	return strings.Join([]string{"SyncSecurityGroupPoliciesResponse", string(data)}, " ")
}
