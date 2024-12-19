package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGdgwRouteTablesRequest Request Object
type ListGdgwRouteTablesRequest struct {

	// 全域接入网关ID
	GdgwId string `json:"gdgw_id"`

	// 地址簇
	AddressFamily *[]ListGdgwRouteTablesRequestAddressFamily `json:"address_family,omitempty"`

	// 下一跳ID
	Nexthop *[]string `json:"nexthop,omitempty"`

	// 目的地址
	Destination *[]string `json:"destination,omitempty"`
}

func (o ListGdgwRouteTablesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGdgwRouteTablesRequest struct{}"
	}

	return strings.Join([]string{"ListGdgwRouteTablesRequest", string(data)}, " ")
}

type ListGdgwRouteTablesRequestAddressFamily struct {
	value string
}

type ListGdgwRouteTablesRequestAddressFamilyEnum struct {
	IPV4 ListGdgwRouteTablesRequestAddressFamily
	IPV6 ListGdgwRouteTablesRequestAddressFamily
}

func GetListGdgwRouteTablesRequestAddressFamilyEnum() ListGdgwRouteTablesRequestAddressFamilyEnum {
	return ListGdgwRouteTablesRequestAddressFamilyEnum{
		IPV4: ListGdgwRouteTablesRequestAddressFamily{
			value: "ipv4",
		},
		IPV6: ListGdgwRouteTablesRequestAddressFamily{
			value: "ipv6",
		},
	}
}

func (c ListGdgwRouteTablesRequestAddressFamily) Value() string {
	return c.value
}

func (c ListGdgwRouteTablesRequestAddressFamily) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGdgwRouteTablesRequestAddressFamily) UnmarshalJSON(b []byte) error {
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
