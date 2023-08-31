package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type EcnErItem struct {

	// 企业连接网络关联企业路由器ID
	Id string `json:"id"`

	// 企业路由器ID
	ErId string `json:"er_id"`

	// 区域ID
	RegionId string `json:"region_id"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 健康状态
	HealthStatus *EcnErItemHealthStatus `json:"health_status,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
}

func (o EcnErItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EcnErItem struct{}"
	}

	return strings.Join([]string{"EcnErItem", string(data)}, " ")
}

type EcnErItemHealthStatus struct {
	value string
}

type EcnErItemHealthStatusEnum struct {
	NORMAL     EcnErItemHealthStatus
	SUB_HEALTH EcnErItemHealthStatus
	FAULT      EcnErItemHealthStatus
}

func GetEcnErItemHealthStatusEnum() EcnErItemHealthStatusEnum {
	return EcnErItemHealthStatusEnum{
		NORMAL: EcnErItemHealthStatus{
			value: "normal",
		},
		SUB_HEALTH: EcnErItemHealthStatus{
			value: "sub_health",
		},
		FAULT: EcnErItemHealthStatus{
			value: "fault",
		},
	}
}

func (c EcnErItemHealthStatus) Value() string {
	return c.value
}

func (c EcnErItemHealthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EcnErItemHealthStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
