package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutes struct {

	// 目的地址
	Destination string `json:"destination"`

	// 下一跳
	Nexthop string `json:"nexthop"`

	// 类型
	Type UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesType `json:"type"`
}

func (o UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutes struct{}"
	}

	return strings.Join([]string{"UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutes", string(data)}, " ")
}

type UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesType struct {
	value string
}

type UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesTypeEnum struct {
	VIF_PEER  UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesType
	PEER_LINK UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesType
	GDGW      UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesType
}

func GetUpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesTypeEnum() UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesTypeEnum {
	return UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesTypeEnum{
		VIF_PEER: UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesType{
			value: "vif_peer",
		},
		PEER_LINK: UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesType{
			value: "peer_link",
		},
		GDGW: UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesType{
			value: "gdgw",
		},
	}
}

func (c UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesType) Value() string {
	return c.value
}

func (c UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateGdgwRoutetableRequestBodyGdgwRoutetableDelRoutesType) UnmarshalJSON(b []byte) error {
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
