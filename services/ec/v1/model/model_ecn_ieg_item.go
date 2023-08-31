package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type EcnIegItem struct {

	// 企业连接网络绑定智能企业网关ID
	Id string `json:"id"`

	// 智能企业网关ID
	IegId *string `json:"ieg_id,omitempty"`

	// 区域ID
	RegionId string `json:"region_id"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 健康状态
	HealthStatus *EcnIegItemHealthStatus `json:"health_status,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
}

func (o EcnIegItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EcnIegItem struct{}"
	}

	return strings.Join([]string{"EcnIegItem", string(data)}, " ")
}

type EcnIegItemHealthStatus struct {
	value string
}

type EcnIegItemHealthStatusEnum struct {
	NORMAL     EcnIegItemHealthStatus
	SUB_HEALTH EcnIegItemHealthStatus
	FAULT      EcnIegItemHealthStatus
}

func GetEcnIegItemHealthStatusEnum() EcnIegItemHealthStatusEnum {
	return EcnIegItemHealthStatusEnum{
		NORMAL: EcnIegItemHealthStatus{
			value: "normal",
		},
		SUB_HEALTH: EcnIegItemHealthStatus{
			value: "sub_health",
		},
		FAULT: EcnIegItemHealthStatus{
			value: "fault",
		},
	}
}

func (c EcnIegItemHealthStatus) Value() string {
	return c.value
}

func (c EcnIegItemHealthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EcnIegItemHealthStatus) UnmarshalJSON(b []byte) error {
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
