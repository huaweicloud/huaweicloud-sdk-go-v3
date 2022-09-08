package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AclApiBindingCreate struct {

	// ACL策略编号
	AclId string `json:"acl_id"`

	// API发布记录编号
	PublishIds []string `json:"publish_ids"`
}

func (o AclApiBindingCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclApiBindingCreate struct{}"
	}

	return strings.Join([]string{"AclApiBindingCreate", string(data)}, " ")
}
