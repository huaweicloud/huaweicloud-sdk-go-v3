package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CloudConnectionQuota 配额实例
type CloudConnectionQuota struct {

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	// 云连接实例ID。
	CloudConnectionId string `json:"cloud_connection_id"`

	// RegionID。
	RegionId string `json:"region_id"`

	// 配额类型： - cloud_connection: 可加载的云连接实例数 - cloud_connection_region: 某云连接实例下可加载的Region数 - cloud_connection_route: 某云连接实例下可加载的路由数 - region_network_instance: 某云连接实例下某个Region下可加载的网络实例数
	QuotaType *CloudConnectionQuotaQuotaType `json:"quota_type,omitempty"`

	// 配额数量。
	QuotaNumber *int32 `json:"quota_number,omitempty"`

	// 配额使用数量。
	QuotaUsed *int32 `json:"quota_used,omitempty"`
}

func (o CloudConnectionQuota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudConnectionQuota struct{}"
	}

	return strings.Join([]string{"CloudConnectionQuota", string(data)}, " ")
}

type CloudConnectionQuotaQuotaType struct {
	value string
}

type CloudConnectionQuotaQuotaTypeEnum struct {
	CLOUD_CONNECTION        CloudConnectionQuotaQuotaType
	CLOUD_CONNECTION_REGION CloudConnectionQuotaQuotaType
	CLOUD_CONNECTION_ROUTE  CloudConnectionQuotaQuotaType
	REGION_NETWORK_INSTANCE CloudConnectionQuotaQuotaType
}

func GetCloudConnectionQuotaQuotaTypeEnum() CloudConnectionQuotaQuotaTypeEnum {
	return CloudConnectionQuotaQuotaTypeEnum{
		CLOUD_CONNECTION: CloudConnectionQuotaQuotaType{
			value: "cloud_connection",
		},
		CLOUD_CONNECTION_REGION: CloudConnectionQuotaQuotaType{
			value: "cloud_connection_region",
		},
		CLOUD_CONNECTION_ROUTE: CloudConnectionQuotaQuotaType{
			value: "cloud_connection_route",
		},
		REGION_NETWORK_INSTANCE: CloudConnectionQuotaQuotaType{
			value: "region_network_instance",
		},
	}
}

func (c CloudConnectionQuotaQuotaType) Value() string {
	return c.value
}

func (c CloudConnectionQuotaQuotaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudConnectionQuotaQuotaType) UnmarshalJSON(b []byte) error {
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
