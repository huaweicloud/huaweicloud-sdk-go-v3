package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegionInstanceInfo 实例信息。
type RegionInstanceInfo struct {

	// 实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例项目id。
	ProjectId *string `json:"project_id,omitempty"`

	// 实例项目名称。
	ProjectName *string `json:"project_name,omitempty"`

	// regionCode编码。
	RegionCode *string `json:"region_code,omitempty"`

	// 数据ip地址列表“,”分割。
	IpAddress *string `json:"ip_address,omitempty"`
}

func (o RegionInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionInstanceInfo struct{}"
	}

	return strings.Join([]string{"RegionInstanceInfo", string(data)}, " ")
}
