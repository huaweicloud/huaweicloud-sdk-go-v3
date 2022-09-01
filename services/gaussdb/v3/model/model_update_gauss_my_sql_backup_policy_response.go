package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateGaussMySqlBackupPolicyResponse struct {

	// 状态信息
	Status *string `json:"status,omitempty" xml:"status"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 实例名称
	InstanceName   *string `json:"instance_name,omitempty" xml:"instance_name"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGaussMySqlBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlBackupPolicyResponse", string(data)}, " ")
}
