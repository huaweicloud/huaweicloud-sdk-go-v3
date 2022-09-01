package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群对象，请参见cluster参数说明
type CdmCreateClusterReqCluster struct {

	// 定时开机的时间，CDM集群会在每天这个时间开机
	ScheduleBootTime *string `json:"scheduleBootTime,omitempty" xml:"scheduleBootTime"`

	// 选择是否启用定时开关机功能。定时开关机功能和自动关机功能不可同时开启
	IsScheduleBootOff *bool `json:"isScheduleBootOff,omitempty" xml:"isScheduleBootOff"`

	// 节点列表，请参见instances参数说明
	Instances *[]Instance `json:"instances,omitempty" xml:"instances"`

	Datastore *Datastore `json:"datastore,omitempty" xml:"datastore"`

	ExtendedProperties *ExtendedProperties `json:"extended_properties,omitempty" xml:"extended_properties"`

	// 定时关机的时间，定时关机时系统不会等待未完成的作业执行完成
	ScheduleOffTime *string `json:"scheduleOffTime,omitempty" xml:"scheduleOffTime"`

	// 指定虚拟私有云ID，用于集群网络配置
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId"`

	// 集群名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 企业项目信息，请参见•sys_tags参数说明
	SysTags *[]SysTags `json:"sys_tags,omitempty" xml:"sys_tags"`

	// 选择是否启用自动关机功能，自动关机功能和定时开关机功能不可同时开启。如果选择自动关机，则当集群中无作业运行且无定时作业时，等待15分钟后集群将自动关机来帮您节约成本
	IsAutoOff *bool `json:"isAutoOff,omitempty" xml:"isAutoOff"`
}

func (o CdmCreateClusterReqCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmCreateClusterReqCluster struct{}"
	}

	return strings.Join([]string{"CdmCreateClusterReqCluster", string(data)}, " ")
}
