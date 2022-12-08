package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 配额实例
type Quota struct {

	// 账号ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 配额类型： - cloud_connection: 可加载的云连接实例数 - cloud_connection_region: 某云连接实例下可加载的Region数 - cloud_connection_route: 某云连接实例下可加载的路由数 - region_network_instance: 某云连接实例下某个Region下可加载的网络实例数
	QuotaType *QuotaQuotaType `json:"quota_type,omitempty"`

	// 配额数量。
	QuotaNumber *int32 `json:"quota_number,omitempty"`

	// 配额使用数量。
	QuotaUsed *int32 `json:"quota_used,omitempty"`

	// 云连接ID。
	CloudConnectionId *string `json:"cloud_connection_id,omitempty"`

	// 网络实例的RegionID。
	RegionId *string `json:"region_id,omitempty"`
}

func (o Quota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quota struct{}"
	}

	return strings.Join([]string{"Quota", string(data)}, " ")
}

type QuotaQuotaType struct {
	value string
}

type QuotaQuotaTypeEnum struct {
	CLOUD_CONNECTION        QuotaQuotaType
	CLOUD_CONNECTION_REGION QuotaQuotaType
	CLOUD_CONNECTION_ROUTE  QuotaQuotaType
	REGION_NETWORK_INSTANCE QuotaQuotaType
}

func GetQuotaQuotaTypeEnum() QuotaQuotaTypeEnum {
	return QuotaQuotaTypeEnum{
		CLOUD_CONNECTION: QuotaQuotaType{
			value: "cloud_connection",
		},
		CLOUD_CONNECTION_REGION: QuotaQuotaType{
			value: "cloud_connection_region",
		},
		CLOUD_CONNECTION_ROUTE: QuotaQuotaType{
			value: "cloud_connection_route",
		},
		REGION_NETWORK_INSTANCE: QuotaQuotaType{
			value: "region_network_instance",
		},
	}
}

func (c QuotaQuotaType) Value() string {
	return c.value
}

func (c QuotaQuotaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QuotaQuotaType) UnmarshalJSON(b []byte) error {
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
