package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterScalingReq struct {

	// 服务ID，为扩展接口，预留此参数。用户不需要配置。
	ServiceId *string `json:"service_id,omitempty" xml:"service_id"`

	// 套餐ID，为扩展接口，预留此参数。用户不需要配置。
	PlanId *string `json:"plan_id,omitempty" xml:"plan_id"`

	Parameters *ClusterScalingParams `json:"parameters" xml:"parameters"`

	// 扩展接口，预留此参数。用户不需要配置。
	PreviousValues map[string]string `json:"previous_values,omitempty" xml:"previous_values"`
}

func (o ClusterScalingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterScalingReq struct{}"
	}

	return strings.Join([]string{"ClusterScalingReq", string(data)}, " ")
}
