package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEffectiveRoutesRequest Request Object
type ListEffectiveRoutesRequest struct {

	// 路由表ID
	RouteTableId string `json:"route_table_id"`

	// 每页返回的个数。 取值范围：0~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条记录的企业路由器实例的id，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 路由目的地址
	Destination *[]string `json:"destination,omitempty"`

	// 连接资源类型:vpc|vpn|vgw|peering
	ResourceType *[]ListEffectiveRoutesRequestResourceType `json:"resource_type,omitempty"`
}

func (o ListEffectiveRoutesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEffectiveRoutesRequest struct{}"
	}

	return strings.Join([]string{"ListEffectiveRoutesRequest", string(data)}, " ")
}

type ListEffectiveRoutesRequestResourceType struct {
	value string
}

type ListEffectiveRoutesRequestResourceTypeEnum struct {
	VPC     ListEffectiveRoutesRequestResourceType
	VPN     ListEffectiveRoutesRequestResourceType
	DGW     ListEffectiveRoutesRequestResourceType
	VGW     ListEffectiveRoutesRequestResourceType
	PEERING ListEffectiveRoutesRequestResourceType
	CAN     ListEffectiveRoutesRequestResourceType
	GDGW    ListEffectiveRoutesRequestResourceType
}

func GetListEffectiveRoutesRequestResourceTypeEnum() ListEffectiveRoutesRequestResourceTypeEnum {
	return ListEffectiveRoutesRequestResourceTypeEnum{
		VPC: ListEffectiveRoutesRequestResourceType{
			value: "vpc",
		},
		VPN: ListEffectiveRoutesRequestResourceType{
			value: "vpn",
		},
		DGW: ListEffectiveRoutesRequestResourceType{
			value: "dgw",
		},
		VGW: ListEffectiveRoutesRequestResourceType{
			value: "vgw",
		},
		PEERING: ListEffectiveRoutesRequestResourceType{
			value: "peering",
		},
		CAN: ListEffectiveRoutesRequestResourceType{
			value: "can",
		},
		GDGW: ListEffectiveRoutesRequestResourceType{
			value: "gdgw",
		},
	}
}

func (c ListEffectiveRoutesRequestResourceType) Value() string {
	return c.value
}

func (c ListEffectiveRoutesRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEffectiveRoutesRequestResourceType) UnmarshalJSON(b []byte) error {
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
