package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupPoliciesRequest Request Object
type ShowBackupPoliciesRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowBackupPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupPoliciesRequest", string(data)}, " ")
}
