package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowEcnWithIegResponse Response Object
type ShowEcnWithIegResponse struct {

	// 企业连接网络绑定智能企业网关ID
	Id *string `json:"id,omitempty"`

	// 智能企业网关ID
	IegId *string `json:"ieg_id,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 健康状态
	HealthStatus *ShowEcnWithIegResponseHealthStatus `json:"health_status,omitempty"`

	// 创建时间
	CreatedAt      *sdktime.SdkTime `json:"created_at,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowEcnWithIegResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEcnWithIegResponse struct{}"
	}

	return strings.Join([]string{"ShowEcnWithIegResponse", string(data)}, " ")
}

type ShowEcnWithIegResponseHealthStatus struct {
	value string
}

type ShowEcnWithIegResponseHealthStatusEnum struct {
	NORMAL     ShowEcnWithIegResponseHealthStatus
	SUB_HEALTH ShowEcnWithIegResponseHealthStatus
	FAULT      ShowEcnWithIegResponseHealthStatus
}

func GetShowEcnWithIegResponseHealthStatusEnum() ShowEcnWithIegResponseHealthStatusEnum {
	return ShowEcnWithIegResponseHealthStatusEnum{
		NORMAL: ShowEcnWithIegResponseHealthStatus{
			value: "normal",
		},
		SUB_HEALTH: ShowEcnWithIegResponseHealthStatus{
			value: "sub_health",
		},
		FAULT: ShowEcnWithIegResponseHealthStatus{
			value: "fault",
		},
	}
}

func (c ShowEcnWithIegResponseHealthStatus) Value() string {
	return c.value
}

func (c ShowEcnWithIegResponseHealthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEcnWithIegResponseHealthStatus) UnmarshalJSON(b []byte) error {
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
