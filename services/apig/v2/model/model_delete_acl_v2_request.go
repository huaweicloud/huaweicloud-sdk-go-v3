package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAclV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 需要删除的ACL ID

	AclId string `json:"acl_id"`
}

func (o DeleteAclV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAclV2Request struct{}"
	}

	return strings.Join([]string{"DeleteAclV2Request", string(data)}, " ")
}
