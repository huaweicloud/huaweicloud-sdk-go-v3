package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceInfoForMetric 实例信息
type InstanceInfoForMetric struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o InstanceInfoForMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceInfoForMetric struct{}"
	}

	return strings.Join([]string{"InstanceInfoForMetric", string(data)}, " ")
}
