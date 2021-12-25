package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AclBatchDelete struct {
	// 需要删除的ACL策略ID列表

	Acls *[]string `json:"acls,omitempty"`
}

func (o AclBatchDelete) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclBatchDelete struct{}"
	}

	return strings.Join([]string{"AclBatchDelete", string(data)}, " ")
}
