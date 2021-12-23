package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 负载均衡器响应体
type LoadbalancerResp struct {
	// 负载均衡器ID

	Id string `json:"id"`
	// 负载均衡器所在的项目ID。

	TenantId string `json:"tenant_id"`
	// 负载均衡器名称。

	Name string `json:"name"`
	// 负载均衡器的描述信息

	Description string `json:"description"`
	// 负载均衡器所在的子网ID。仅支持内网类型。

	VipSubnetId string `json:"vip_subnet_id"`
	// 负载均衡器虚拟IP对应的端口ID

	VipPortId string `json:"vip_port_id"`
	// 负载均衡器的虚拟IP。

	VipAddress string `json:"vip_address"`
	// 负载均衡器关联的监听器ID的列表

	Listeners []ResourceList `json:"listeners"`
	// 负载均衡器关联的后端云服务器组ID的列表。

	Pools []ResourceList `json:"pools"`
	// 负载均衡器的供应者名称。只支持vlb

	Provider string `json:"provider"`
	// 负载均衡器的操作状态

	OperatingStatus LoadbalancerRespOperatingStatus `json:"operating_status"`
	// 负载均衡器的配置状态

	ProvisioningStatus LoadbalancerRespProvisioningStatus `json:"provisioning_status"`
	// 负载均衡器的管理状态。只支持设定为true，该字段的值无实际意义。

	AdminStateUp bool `json:"admin_state_up"`
	// 负载均衡器的创建时间

	CreatedAt string `json:"created_at"`
	// 负载均衡器的更新时间

	UpdatedAt string `json:"updated_at"`
	// 负载均衡器的企业项目ID。

	EnterpriseProjectId string `json:"enterprise_project_id"`
	// 负载均衡器所在的项目ID。

	ProjectId string `json:"project_id"`
	// 负载均衡器的标签列表

	Tags []string `json:"tags"`
}

func (o LoadbalancerResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadbalancerResp struct{}"
	}

	return strings.Join([]string{"LoadbalancerResp", string(data)}, " ")
}

type LoadbalancerRespOperatingStatus struct {
	value string
}

type LoadbalancerRespOperatingStatusEnum struct {
	ONLINE     LoadbalancerRespOperatingStatus
	OFFLINE    LoadbalancerRespOperatingStatus
	DEGRADED   LoadbalancerRespOperatingStatus
	DISABLED   LoadbalancerRespOperatingStatus
	NO_MONITOR LoadbalancerRespOperatingStatus
}

func GetLoadbalancerRespOperatingStatusEnum() LoadbalancerRespOperatingStatusEnum {
	return LoadbalancerRespOperatingStatusEnum{
		ONLINE: LoadbalancerRespOperatingStatus{
			value: "ONLINE",
		},
		OFFLINE: LoadbalancerRespOperatingStatus{
			value: "OFFLINE",
		},
		DEGRADED: LoadbalancerRespOperatingStatus{
			value: "DEGRADED",
		},
		DISABLED: LoadbalancerRespOperatingStatus{
			value: "DISABLED",
		},
		NO_MONITOR: LoadbalancerRespOperatingStatus{
			value: "NO_MONITOR",
		},
	}
}

func (c LoadbalancerRespOperatingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadbalancerRespOperatingStatus) UnmarshalJSON(b []byte) error {
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

type LoadbalancerRespProvisioningStatus struct {
	value string
}

type LoadbalancerRespProvisioningStatusEnum struct {
	ACTIVE         LoadbalancerRespProvisioningStatus
	PENDING_CREATE LoadbalancerRespProvisioningStatus
	ERROR          LoadbalancerRespProvisioningStatus
}

func GetLoadbalancerRespProvisioningStatusEnum() LoadbalancerRespProvisioningStatusEnum {
	return LoadbalancerRespProvisioningStatusEnum{
		ACTIVE: LoadbalancerRespProvisioningStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: LoadbalancerRespProvisioningStatus{
			value: "PENDING_CREATE",
		},
		ERROR: LoadbalancerRespProvisioningStatus{
			value: "ERROR",
		},
	}
}

func (c LoadbalancerRespProvisioningStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadbalancerRespProvisioningStatus) UnmarshalJSON(b []byte) error {
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
