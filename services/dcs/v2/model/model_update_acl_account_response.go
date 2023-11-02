package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAclAccountResponse Response Object
type UpdateAclAccountResponse struct {

	// 账号所属实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 修改ACL账号访问权限JOB的ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAclAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclAccountResponse struct{}"
	}

	return strings.Join([]string{"UpdateAclAccountResponse", string(data)}, " ")
}
