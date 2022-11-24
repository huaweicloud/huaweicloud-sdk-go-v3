package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// 终端节点实例。
type EndpointDetail struct {

	// 终端节点ID。
	Id *string `json:"id,omitempty"`

	// 对应后端资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 终端节点组ID。
	EndpointGroupId *string `json:"endpoint_group_id,omitempty"`

	ResourceType *EndpointType `json:"resource_type,omitempty"`

	Status *ConfigStatus `json:"status,omitempty"`

	// 终端节点权重。
	Weight *int32 `json:"weight,omitempty"`

	// 终端的健康状态，取值： - INITIAL：初始 - HEALTHY：正常 - UNHEALTHY：异常 - NO_MONITOR：未监控
	HealthState *EndpointDetailHealthState `json:"health_state,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 租户ID。
	DomainId *string `json:"domain_id,omitempty"`

	// IP地址。
	IpAddress *string `json:"ip_address,omitempty"`

	FrozenInfo *FrozenInfo `json:"frozen_info,omitempty"`
}

func (o EndpointDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointDetail struct{}"
	}

	return strings.Join([]string{"EndpointDetail", string(data)}, " ")
}

type EndpointDetailHealthState struct {
	value string
}

type EndpointDetailHealthStateEnum struct {
	INITIAL    EndpointDetailHealthState
	HEALTHY    EndpointDetailHealthState
	UNHEALTHY  EndpointDetailHealthState
	NO_MONITOR EndpointDetailHealthState
}

func GetEndpointDetailHealthStateEnum() EndpointDetailHealthStateEnum {
	return EndpointDetailHealthStateEnum{
		INITIAL: EndpointDetailHealthState{
			value: "INITIAL",
		},
		HEALTHY: EndpointDetailHealthState{
			value: "HEALTHY",
		},
		UNHEALTHY: EndpointDetailHealthState{
			value: "UNHEALTHY",
		},
		NO_MONITOR: EndpointDetailHealthState{
			value: "NO_MONITOR",
		},
	}
}

func (c EndpointDetailHealthState) Value() string {
	return c.value
}

func (c EndpointDetailHealthState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointDetailHealthState) UnmarshalJSON(b []byte) error {
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
