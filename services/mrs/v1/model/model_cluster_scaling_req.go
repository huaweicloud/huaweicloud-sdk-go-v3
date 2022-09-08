package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterScalingReq struct {

	// 服务ID，为扩展接口，预留此参数。用户不需要配置。
	ServiceId *string `json:"service_id,omitempty"`

	// 套餐ID，为扩展接口，预留此参数。用户不需要配置。
	PlanId *string `json:"plan_id,omitempty"`

	Parameters *ClusterScalingParams `json:"parameters"`

	// 扩展接口，预留此参数。用户不需要配置。
	PreviousValues map[string]string `json:"previous_values,omitempty"`
}

func (o ClusterScalingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterScalingReq struct{}"
	}

	return strings.Join([]string{"ClusterScalingReq", string(data)}, " ")
}
