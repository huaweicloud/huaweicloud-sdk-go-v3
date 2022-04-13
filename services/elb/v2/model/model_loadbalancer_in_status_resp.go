package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 负载均衡器对象，用于负载均衡器状态树中
type LoadbalancerInStatusResp struct {
	// 负载均衡器名称

	Name string `json:"name"`
	// 负载均衡器ID

	Id string `json:"id"`
	// 负载均衡器关联的监听器列表

	Listeners []ListenersInStatusResp `json:"listeners"`
	// 负载均衡器关联的后端云服务器组列表

	Pools []PoolsInStatusResp `json:"pools"`
	// 负载均衡器的操作状态；该字段为预留字段，暂未启用。默认为ONLINE。

	OperatingStatus LoadbalancerInStatusRespOperatingStatus `json:"operating_status"`
	// 负载均衡器的配置状态；该字段为预留字段，暂未启用。默认为ACTIVE。

	ProvisioningStatus LoadbalancerInStatusRespProvisioningStatus `json:"provisioning_status"`
}

func (o LoadbalancerInStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadbalancerInStatusResp struct{}"
	}

	return strings.Join([]string{"LoadbalancerInStatusResp", string(data)}, " ")
}

type LoadbalancerInStatusRespOperatingStatus struct {
	value string
}

type LoadbalancerInStatusRespOperatingStatusEnum struct {
	ONLINE     LoadbalancerInStatusRespOperatingStatus
	OFFLINE    LoadbalancerInStatusRespOperatingStatus
	DEGRADED   LoadbalancerInStatusRespOperatingStatus
	DISABLED   LoadbalancerInStatusRespOperatingStatus
	NO_MONITOR LoadbalancerInStatusRespOperatingStatus
}

func GetLoadbalancerInStatusRespOperatingStatusEnum() LoadbalancerInStatusRespOperatingStatusEnum {
	return LoadbalancerInStatusRespOperatingStatusEnum{
		ONLINE: LoadbalancerInStatusRespOperatingStatus{
			value: "ONLINE",
		},
		OFFLINE: LoadbalancerInStatusRespOperatingStatus{
			value: "OFFLINE",
		},
		DEGRADED: LoadbalancerInStatusRespOperatingStatus{
			value: "DEGRADED",
		},
		DISABLED: LoadbalancerInStatusRespOperatingStatus{
			value: "DISABLED",
		},
		NO_MONITOR: LoadbalancerInStatusRespOperatingStatus{
			value: "NO_MONITOR",
		},
	}
}

func (c LoadbalancerInStatusRespOperatingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadbalancerInStatusRespOperatingStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type LoadbalancerInStatusRespProvisioningStatus struct {
	value string
}

type LoadbalancerInStatusRespProvisioningStatusEnum struct {
	ACTIVE         LoadbalancerInStatusRespProvisioningStatus
	PENDING_CREATE LoadbalancerInStatusRespProvisioningStatus
	ERROR          LoadbalancerInStatusRespProvisioningStatus
}

func GetLoadbalancerInStatusRespProvisioningStatusEnum() LoadbalancerInStatusRespProvisioningStatusEnum {
	return LoadbalancerInStatusRespProvisioningStatusEnum{
		ACTIVE: LoadbalancerInStatusRespProvisioningStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: LoadbalancerInStatusRespProvisioningStatus{
			value: "PENDING_CREATE",
		},
		ERROR: LoadbalancerInStatusRespProvisioningStatus{
			value: "ERROR",
		},
	}
}

func (c LoadbalancerInStatusRespProvisioningStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadbalancerInStatusRespProvisioningStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
