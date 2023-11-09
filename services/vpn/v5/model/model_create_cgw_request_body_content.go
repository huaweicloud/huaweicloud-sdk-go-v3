package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateCgwRequestBodyContent struct {

	// 网关名称
	Name *string `json:"name,omitempty"`

	// 网关路由模式
	RouteMode *CreateCgwRequestBodyContentRouteMode `json:"route_mode,omitempty"`

	// 网关的bgp asn号，仅当route_mode为bgp时需要，默认值为65000
	BgpAsn *int64 `json:"bgp_asn,omitempty"`

	// 网关ip地址
	Ip string `json:"ip"`

	CaCertificate *CaCertificateRequest `json:"ca_certificate,omitempty"`
}

func (o CreateCgwRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCgwRequestBodyContent struct{}"
	}

	return strings.Join([]string{"CreateCgwRequestBodyContent", string(data)}, " ")
}

type CreateCgwRequestBodyContentRouteMode struct {
	value string
}

type CreateCgwRequestBodyContentRouteModeEnum struct {
	STATIC CreateCgwRequestBodyContentRouteMode
	BGP    CreateCgwRequestBodyContentRouteMode
}

func GetCreateCgwRequestBodyContentRouteModeEnum() CreateCgwRequestBodyContentRouteModeEnum {
	return CreateCgwRequestBodyContentRouteModeEnum{
		STATIC: CreateCgwRequestBodyContentRouteMode{
			value: "static",
		},
		BGP: CreateCgwRequestBodyContentRouteMode{
			value: "bgp",
		},
	}
}

func (c CreateCgwRequestBodyContentRouteMode) Value() string {
	return c.value
}

func (c CreateCgwRequestBodyContentRouteMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateCgwRequestBodyContentRouteMode) UnmarshalJSON(b []byte) error {
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
