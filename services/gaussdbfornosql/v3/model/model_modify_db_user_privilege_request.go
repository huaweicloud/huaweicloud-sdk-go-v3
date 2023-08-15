package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyDbUserPrivilegeRequest Request Object
type ModifyDbUserPrivilegeRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *RedisModifyDbUserPrivilegeRequest `json:"body,omitempty"`
}

func (o ModifyDbUserPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDbUserPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"ModifyDbUserPrivilegeRequest", string(data)}, " ")
}
