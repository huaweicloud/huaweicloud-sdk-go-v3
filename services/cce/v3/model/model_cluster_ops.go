package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterOps **参数解释：** 集群运维参数配置。 **约束限制：** 不涉及
type ClusterOps struct {
	Alarm *AlarmInfo `json:"alarm"`
}

func (o ClusterOps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterOps struct{}"
	}

	return strings.Join([]string{"ClusterOps", string(data)}, " ")
}
