package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopLogAutoBackupPolicyRequest struct {
	// 指定待关闭的集群ID。

	ClusterId string `json:"cluster_id"`
}

func (o StopLogAutoBackupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopLogAutoBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"StopLogAutoBackupPolicyRequest", string(data)}, " ")
}
