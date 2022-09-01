package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAclStrategyV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// ACL策略的编号
	AclId string `json:"acl_id" xml:"acl_id"`

	Body *ApiAclCreate `json:"body,omitempty" xml:"body"`
}

func (o UpdateAclStrategyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclStrategyV2Request struct{}"
	}

	return strings.Join([]string{"UpdateAclStrategyV2Request", string(data)}, " ")
}
