package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ResponseCustomerGateway struct {

	// 网关的ID
	Id *string `json:"id,omitempty"`

	// 网关名称
	Name *string `json:"name,omitempty"`

	// 网关路由模式
	RouteMode *ResponseCustomerGatewayRouteMode `json:"route_mode,omitempty"`

	// 网关的bgp asn号
	BgpAsn *int64 `json:"bgp_asn,omitempty"`

	// 网关ip地址
	Ip *string `json:"ip,omitempty"`

	CaCertificate *CaCertificate `json:"ca_certificate,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ResponseCustomerGateway) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseCustomerGateway struct{}"
	}

	return strings.Join([]string{"ResponseCustomerGateway", string(data)}, " ")
}

type ResponseCustomerGatewayRouteMode struct {
	value string
}

type ResponseCustomerGatewayRouteModeEnum struct {
	STATIC ResponseCustomerGatewayRouteMode
	BGP    ResponseCustomerGatewayRouteMode
}

func GetResponseCustomerGatewayRouteModeEnum() ResponseCustomerGatewayRouteModeEnum {
	return ResponseCustomerGatewayRouteModeEnum{
		STATIC: ResponseCustomerGatewayRouteMode{
			value: "static",
		},
		BGP: ResponseCustomerGatewayRouteMode{
			value: "bgp",
		},
	}
}

func (c ResponseCustomerGatewayRouteMode) Value() string {
	return c.value
}

func (c ResponseCustomerGatewayRouteMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResponseCustomerGatewayRouteMode) UnmarshalJSON(b []byte) error {
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
