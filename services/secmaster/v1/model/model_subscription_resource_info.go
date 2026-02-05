package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SubscriptionResourceInfo 租户资源类
type SubscriptionResourceInfo struct {

	// 资源Id
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源规格
	ResourceSize *int32 `json:"resource_size,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源规格编码
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 创建时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 到期时间戳，只有按需资源有该字段
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// 资源状态，目前返回正常运行的资源，其状态值为0
	ResourceStatus *int32 `json:"resource_status,omitempty"`

	// 订单Id，包周期资源有该字段
	OrderId *string `json:"order_id,omitempty"`

	// 计费模式，目前有包周期（包年包月）PREPAID、按需POSTPAID，大小写不敏感
	ChargingMode *SubscriptionResourceInfoChargingMode `json:"charging_mode,omitempty"`

	// 当前资源是否能进行按需转包周期操作
	ToPeriod *bool `json:"to_period,omitempty"`

	// 资源列表
	TagList *[]TagInfo `json:"tag_list,omitempty"`
}

func (o SubscriptionResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionResourceInfo struct{}"
	}

	return strings.Join([]string{"SubscriptionResourceInfo", string(data)}, " ")
}

type SubscriptionResourceInfoChargingMode struct {
	value string
}

type SubscriptionResourceInfoChargingModeEnum struct {
	PREPAID  SubscriptionResourceInfoChargingMode
	POSTPAID SubscriptionResourceInfoChargingMode
}

func GetSubscriptionResourceInfoChargingModeEnum() SubscriptionResourceInfoChargingModeEnum {
	return SubscriptionResourceInfoChargingModeEnum{
		PREPAID: SubscriptionResourceInfoChargingMode{
			value: "PREPAID",
		},
		POSTPAID: SubscriptionResourceInfoChargingMode{
			value: "POSTPAID",
		},
	}
}

func (c SubscriptionResourceInfoChargingMode) Value() string {
	return c.value
}

func (c SubscriptionResourceInfoChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubscriptionResourceInfoChargingMode) UnmarshalJSON(b []byte) error {
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
