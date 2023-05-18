package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 容灾可用集群信息
type DisasterRecoveryClusterVo struct {

	// 集群ID
	Id *string `json:"id,omitempty"`

	// 集群名称
	Name *string `json:"name,omitempty"`
}

func (o DisasterRecoveryClusterVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisasterRecoveryClusterVo struct{}"
	}

	return strings.Join([]string{"DisasterRecoveryClusterVo", string(data)}, " ")
}
