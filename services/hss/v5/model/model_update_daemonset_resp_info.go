package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDaemonsetRespInfo 升级Ds响应信息
type UpdateDaemonsetRespInfo struct {

	// 失败原因
	FailedReson *string `json:"failed_reson,omitempty"`

	// 集群Id
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o UpdateDaemonsetRespInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDaemonsetRespInfo struct{}"
	}

	return strings.Join([]string{"UpdateDaemonsetRespInfo", string(data)}, " ")
}
