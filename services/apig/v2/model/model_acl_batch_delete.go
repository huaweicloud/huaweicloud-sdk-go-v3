package model

import (
	"encoding/json"

	"strings"
)

type AclBatchDelete struct {
	// 需要删除的ACL策略ID列表

	Acls *[]string `json:"acls,omitempty"`
}

func (o AclBatchDelete) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AclBatchDelete struct{}"
	}

	return strings.Join([]string{"AclBatchDelete", string(data)}, " ")
}
