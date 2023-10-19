package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListCloudConnectionQuotasRequest Request Object
type ListCloudConnectionQuotasRequest struct {

	// 每页返回的个数。 取值范围：1~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向翻页。 翻页过程中，查询条件不能修改，包括过滤条件，排序条件，limit。
	Marker *string `json:"marker,omitempty"`

	// 配额类型： - cloud_connection: 可加载的云连接实例数 - cloud_connection_region: 某云连接实例下可加载的Region数 - cloud_connection_route: 某云连接实例下可加载的路由数 - region_network_instance: 某云连接实例下某个Region下可加载的网络实例数
	QuotaType ListCloudConnectionQuotasRequestQuotaType `json:"quota_type"`

	// 云连接ID。当查询cloud_connection_region、cloud_connection_route、region_network_instance三种类型的配额时需要填写此参数。
	CloudConnectionId *string `json:"cloud_connection_id,omitempty"`

	// 区域ID。当查询region_network_instance类型的配额时需要填写此参数。
	RegionId *string `json:"region_id,omitempty"`
}

func (o ListCloudConnectionQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionQuotasRequest struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionQuotasRequest", string(data)}, " ")
}

type ListCloudConnectionQuotasRequestQuotaType struct {
	value string
}

type ListCloudConnectionQuotasRequestQuotaTypeEnum struct {
	CLOUD_CONNECTION        ListCloudConnectionQuotasRequestQuotaType
	CLOUD_CONNECTION_REGION ListCloudConnectionQuotasRequestQuotaType
	CLOUD_CONNECTION_ROUTE  ListCloudConnectionQuotasRequestQuotaType
	REGION_NETWORK_INSTANCE ListCloudConnectionQuotasRequestQuotaType
}

func GetListCloudConnectionQuotasRequestQuotaTypeEnum() ListCloudConnectionQuotasRequestQuotaTypeEnum {
	return ListCloudConnectionQuotasRequestQuotaTypeEnum{
		CLOUD_CONNECTION: ListCloudConnectionQuotasRequestQuotaType{
			value: "cloud_connection",
		},
		CLOUD_CONNECTION_REGION: ListCloudConnectionQuotasRequestQuotaType{
			value: "cloud_connection_region",
		},
		CLOUD_CONNECTION_ROUTE: ListCloudConnectionQuotasRequestQuotaType{
			value: "cloud_connection_route",
		},
		REGION_NETWORK_INSTANCE: ListCloudConnectionQuotasRequestQuotaType{
			value: "region_network_instance",
		},
	}
}

func (c ListCloudConnectionQuotasRequestQuotaType) Value() string {
	return c.value
}

func (c ListCloudConnectionQuotasRequestQuotaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListCloudConnectionQuotasRequestQuotaType) UnmarshalJSON(b []byte) error {
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
