package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// AddEcnWithIegResponse Response Object
type AddEcnWithIegResponse struct {

	// 企业连接网络绑定智能企业网关ID
	Id *string `json:"id,omitempty"`

	// 智能企业网关ID
	IegId *string `json:"ieg_id,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 健康状态
	HealthStatus *AddEcnWithIegResponseHealthStatus `json:"health_status,omitempty"`

	// 创建时间
	CreatedAt      *sdktime.SdkTime `json:"created_at,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o AddEcnWithIegResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEcnWithIegResponse struct{}"
	}

	return strings.Join([]string{"AddEcnWithIegResponse", string(data)}, " ")
}

type AddEcnWithIegResponseHealthStatus struct {
	value string
}

type AddEcnWithIegResponseHealthStatusEnum struct {
	NORMAL     AddEcnWithIegResponseHealthStatus
	SUB_HEALTH AddEcnWithIegResponseHealthStatus
	FAULT      AddEcnWithIegResponseHealthStatus
}

func GetAddEcnWithIegResponseHealthStatusEnum() AddEcnWithIegResponseHealthStatusEnum {
	return AddEcnWithIegResponseHealthStatusEnum{
		NORMAL: AddEcnWithIegResponseHealthStatus{
			value: "normal",
		},
		SUB_HEALTH: AddEcnWithIegResponseHealthStatus{
			value: "sub_health",
		},
		FAULT: AddEcnWithIegResponseHealthStatus{
			value: "fault",
		},
	}
}

func (c AddEcnWithIegResponseHealthStatus) Value() string {
	return c.value
}

func (c AddEcnWithIegResponseHealthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddEcnWithIegResponseHealthStatus) UnmarshalJSON(b []byte) error {
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
