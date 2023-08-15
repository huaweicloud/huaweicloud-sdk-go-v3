package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopLogAutoBackupPolicyRequest Request Object
type StopLogAutoBackupPolicyRequest struct {

	// 指定关闭日志自动备份的集群ID。
	ClusterId string `json:"cluster_id"`
}

func (o StopLogAutoBackupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopLogAutoBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"StopLogAutoBackupPolicyRequest", string(data)}, " ")
}
