package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAclV2Request Request Object
type DeleteAclV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// ACL策略的编号
	AclId string `json:"acl_id"`
}

func (o DeleteAclV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAclV2Request struct{}"
	}

	return strings.Join([]string{"DeleteAclV2Request", string(data)}, " ")
}
