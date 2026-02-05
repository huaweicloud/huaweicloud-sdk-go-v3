package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SubscriptionResource 租户资源类
type SubscriptionResource struct {

	// 资源Id
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称
	ResourceTypeName *string `json:"resource_type_name,omitempty"`

	// 资源规格
	ResourceSize *int32 `json:"resource_size,omitempty"`

	// 订单来源，默认数据为SecMaster
	CloudService *string `json:"cloud_service,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 资源规格编码
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 当前资源是否能进行按需转包周期操作
	ToPeriod *bool `json:"to_period,omitempty"`

	// 创建时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 到期时间戳，只有包年包月资源才有该字段
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// 资源状态，目前返回正常运行的资源，其状态值为0
	ResourceStatus *int32 `json:"resource_status,omitempty"`

	// 订单Id，包周期资源有该字段
	OrderId *string `json:"order_id,omitempty"`

	// 计费模式，目前有包周期（包年包月）PREPAID、按需POSTPAID，大小写不敏感
	ChargingMode *SubscriptionResourceChargingMode `json:"charging_mode,omitempty"`

	// 资源列表
	TagList *[]SubscriptionTagInfo `json:"tag_list,omitempty"`

	// 资源使用量，当usage=true时返回该字段
	Usages *[]ResourceUsage `json:"usages,omitempty"`
}

func (o SubscriptionResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionResource struct{}"
	}

	return strings.Join([]string{"SubscriptionResource", string(data)}, " ")
}

type SubscriptionResourceChargingMode struct {
	value string
}

type SubscriptionResourceChargingModeEnum struct {
	PREPAID  SubscriptionResourceChargingMode
	POSTPAID SubscriptionResourceChargingMode
}

func GetSubscriptionResourceChargingModeEnum() SubscriptionResourceChargingModeEnum {
	return SubscriptionResourceChargingModeEnum{
		PREPAID: SubscriptionResourceChargingMode{
			value: "PREPAID",
		},
		POSTPAID: SubscriptionResourceChargingMode{
			value: "POSTPAID",
		},
	}
}

func (c SubscriptionResourceChargingMode) Value() string {
	return c.value
}

func (c SubscriptionResourceChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubscriptionResourceChargingMode) UnmarshalJSON(b []byte) error {
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
