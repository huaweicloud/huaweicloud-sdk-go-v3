package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailsOfAclPolicyV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// ACL策略的编号

	AclId string `json:"acl_id"`
}

func (o ShowDetailsOfAclPolicyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfAclPolicyV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfAclPolicyV2Request", string(data)}, " ")
}
