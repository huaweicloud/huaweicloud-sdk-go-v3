package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RouteTypeOfGdgw 下一跳类型: - vif_peer: 虚拟接口对等体 - gdgw: 全域接入网关
type RouteTypeOfGdgw struct {
	value string
}

type RouteTypeOfGdgwEnum struct {
	VIF_PEER RouteTypeOfGdgw
	GDGW     RouteTypeOfGdgw
}

func GetRouteTypeOfGdgwEnum() RouteTypeOfGdgwEnum {
	return RouteTypeOfGdgwEnum{
		VIF_PEER: RouteTypeOfGdgw{
			value: "vif_peer",
		},
		GDGW: RouteTypeOfGdgw{
			value: "gdgw",
		},
	}
}

func (c RouteTypeOfGdgw) Value() string {
	return c.value
}

func (c RouteTypeOfGdgw) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RouteTypeOfGdgw) UnmarshalJSON(b []byte) error {
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
