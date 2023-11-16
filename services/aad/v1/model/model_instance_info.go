package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceInfo struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例IP
	Ips *[]InstanceIpInfo `json:"ips,omitempty"`

	// 实例过期时间
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// 业务带宽
	ServiceBandwidth *int32 `json:"service_bandwidth,omitempty"`

	// 实例状态
	InstanceStatus *int32 `json:"instance_status,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o InstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceInfo struct{}"
	}

	return strings.Join([]string{"InstanceInfo", string(data)}, " ")
}
