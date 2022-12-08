package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListQuotasRequest struct {

	// 分页查询时，每页返回的个数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询时，上一页最后一条记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 配额类型： - cloud_connection: 可加载的云连接实例数 - cloud_connection_region: 某云连接实例下可加载的Region数 - cloud_connection_route: 某云连接实例下可加载的路由数 - region_network_instance: 某云连接实例下某个Region下可加载的网络实例数
	QuotaType *ListQuotasRequestQuotaType `json:"quota_type,omitempty"`
}

func (o ListQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotasRequest struct{}"
	}

	return strings.Join([]string{"ListQuotasRequest", string(data)}, " ")
}

type ListQuotasRequestQuotaType struct {
	value string
}

type ListQuotasRequestQuotaTypeEnum struct {
	CLOUD_CONNECTION        ListQuotasRequestQuotaType
	CLOUD_CONNECTION_REGION ListQuotasRequestQuotaType
	CLOUD_CONNECTION_ROUTE  ListQuotasRequestQuotaType
	REGION_NETWORK_INSTANCE ListQuotasRequestQuotaType
}

func GetListQuotasRequestQuotaTypeEnum() ListQuotasRequestQuotaTypeEnum {
	return ListQuotasRequestQuotaTypeEnum{
		CLOUD_CONNECTION: ListQuotasRequestQuotaType{
			value: "cloud_connection",
		},
		CLOUD_CONNECTION_REGION: ListQuotasRequestQuotaType{
			value: "cloud_connection_region",
		},
		CLOUD_CONNECTION_ROUTE: ListQuotasRequestQuotaType{
			value: "cloud_connection_route",
		},
		REGION_NETWORK_INSTANCE: ListQuotasRequestQuotaType{
			value: "region_network_instance",
		},
	}
}

func (c ListQuotasRequestQuotaType) Value() string {
	return c.value
}

func (c ListQuotasRequestQuotaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListQuotasRequestQuotaType) UnmarshalJSON(b []byte) error {
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
