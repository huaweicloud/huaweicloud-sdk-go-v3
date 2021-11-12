package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAclStrategyV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// ACL策略的编号。可通过查询ACL信息获取该ID。

	AclId string `json:"acl_id"`

	Body *ApiAclCreate `json:"body,omitempty"`
}

func (o UpdateAclStrategyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclStrategyV2Request struct{}"
	}

	return strings.Join([]string{"UpdateAclStrategyV2Request", string(data)}, " ")
}
