package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// AddEcnWithErResponse Response Object
type AddEcnWithErResponse struct {

	// 企业连接网络关联企业路由器ID
	Id *string `json:"id,omitempty"`

	// 企业路由器ID
	ErId *string `json:"er_id,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 健康状态
	HealthStatus *AddEcnWithErResponseHealthStatus `json:"health_status,omitempty"`

	// 创建时间
	CreatedAt      *sdktime.SdkTime `json:"created_at,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o AddEcnWithErResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEcnWithErResponse struct{}"
	}

	return strings.Join([]string{"AddEcnWithErResponse", string(data)}, " ")
}

type AddEcnWithErResponseHealthStatus struct {
	value string
}

type AddEcnWithErResponseHealthStatusEnum struct {
	NORMAL     AddEcnWithErResponseHealthStatus
	SUB_HEALTH AddEcnWithErResponseHealthStatus
	FAULT      AddEcnWithErResponseHealthStatus
}

func GetAddEcnWithErResponseHealthStatusEnum() AddEcnWithErResponseHealthStatusEnum {
	return AddEcnWithErResponseHealthStatusEnum{
		NORMAL: AddEcnWithErResponseHealthStatus{
			value: "normal",
		},
		SUB_HEALTH: AddEcnWithErResponseHealthStatus{
			value: "sub_health",
		},
		FAULT: AddEcnWithErResponseHealthStatus{
			value: "fault",
		},
	}
}

func (c AddEcnWithErResponseHealthStatus) Value() string {
	return c.value
}

func (c AddEcnWithErResponseHealthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddEcnWithErResponseHealthStatus) UnmarshalJSON(b []byte) error {
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
