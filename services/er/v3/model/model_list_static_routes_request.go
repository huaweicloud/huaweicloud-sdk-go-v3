package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListStaticRoutesRequest struct {

	// 路由表ID
	RouteTableId string `json:"route_table_id"`

	// 每页返回的个数。 取值范围：0~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条记录的企业路由器实例的id，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 路由目的地址
	Destination *[]string `json:"destination,omitempty"`

	// 连接ID
	AttachmentId *[]string `json:"attachment_id,omitempty"`

	// 连接资源类型:vpc|vpn|vgw|peering|can|gdgw
	ResourceType *[]ListStaticRoutesRequestResourceType `json:"resource_type,omitempty"`

	// 按关键字排序，默认按照id排序，可选值:id|name|state
	SortKey *[]string `json:"sort_key,omitempty"`

	// 返回结果按照升序或降序排列，默认为asc,降序为desc
	SortDir *[]ListStaticRoutesRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListStaticRoutesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStaticRoutesRequest struct{}"
	}

	return strings.Join([]string{"ListStaticRoutesRequest", string(data)}, " ")
}

type ListStaticRoutesRequestResourceType struct {
	value string
}

type ListStaticRoutesRequestResourceTypeEnum struct {
	VPC     ListStaticRoutesRequestResourceType
	VPN     ListStaticRoutesRequestResourceType
	DGW     ListStaticRoutesRequestResourceType
	VGW     ListStaticRoutesRequestResourceType
	PEERING ListStaticRoutesRequestResourceType
	CAN     ListStaticRoutesRequestResourceType
	GDGW    ListStaticRoutesRequestResourceType
}

func GetListStaticRoutesRequestResourceTypeEnum() ListStaticRoutesRequestResourceTypeEnum {
	return ListStaticRoutesRequestResourceTypeEnum{
		VPC: ListStaticRoutesRequestResourceType{
			value: "vpc",
		},
		VPN: ListStaticRoutesRequestResourceType{
			value: "vpn",
		},
		DGW: ListStaticRoutesRequestResourceType{
			value: "dgw",
		},
		VGW: ListStaticRoutesRequestResourceType{
			value: "vgw",
		},
		PEERING: ListStaticRoutesRequestResourceType{
			value: "peering",
		},
		CAN: ListStaticRoutesRequestResourceType{
			value: "can",
		},
		GDGW: ListStaticRoutesRequestResourceType{
			value: "gdgw",
		},
	}
}

func (c ListStaticRoutesRequestResourceType) Value() string {
	return c.value
}

func (c ListStaticRoutesRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListStaticRoutesRequestResourceType) UnmarshalJSON(b []byte) error {
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

type ListStaticRoutesRequestSortDir struct {
	value string
}

type ListStaticRoutesRequestSortDirEnum struct {
	ASC  ListStaticRoutesRequestSortDir
	DESC ListStaticRoutesRequestSortDir
}

func GetListStaticRoutesRequestSortDirEnum() ListStaticRoutesRequestSortDirEnum {
	return ListStaticRoutesRequestSortDirEnum{
		ASC: ListStaticRoutesRequestSortDir{
			value: "asc",
		},
		DESC: ListStaticRoutesRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListStaticRoutesRequestSortDir) Value() string {
	return c.value
}

func (c ListStaticRoutesRequestSortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListStaticRoutesRequestSortDir) UnmarshalJSON(b []byte) error {
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
