package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloudAlarmDto struct {

	// 云服务名称
	CloudServiceName *string `json:"cloudServiceName,omitempty"`

	// 云服务区域标识
	CloudServiceRegionId *string `json:"cloudServiceRegionId,omitempty"`

	// 云服务站点：默认中国站
	CloudServiceSite *string `json:"cloudServiceSite,omitempty"`

	// 是否开启CloudAlarm配置
	Enable *string `json:"enable,omitempty"`

	// 告警级别
	Level *string `json:"level,omitempty"`

	// 微服务组名称
	MicroServiceGroupName *string `json:"microServiceGroupName,omitempty"`

	// 微服务名称
	MicroServiceName *string `json:"microServiceName,omitempty"`
}

func (o CloudAlarmDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudAlarmDto struct{}"
	}

	return strings.Join([]string{"CloudAlarmDto", string(data)}, " ")
}
