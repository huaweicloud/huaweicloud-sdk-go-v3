package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutes struct {

	// 目的地址
	Destination string `json:"destination"`

	// 下一跳
	Nexthop string `json:"nexthop"`

	// 类型
	Type UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesType `json:"type"`

	// 说明信息
	Description *string `json:"description,omitempty"`
}

func (o UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutes struct{}"
	}

	return strings.Join([]string{"UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutes", string(data)}, " ")
}

type UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesType struct {
	value string
}

type UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesTypeEnum struct {
	VIF_PEER  UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesType
	PEER_LINK UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesType
	GDGW      UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesType
}

func GetUpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesTypeEnum() UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesTypeEnum {
	return UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesTypeEnum{
		VIF_PEER: UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesType{
			value: "vif_peer",
		},
		PEER_LINK: UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesType{
			value: "peer_link",
		},
		GDGW: UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesType{
			value: "gdgw",
		},
	}
}

func (c UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesType) Value() string {
	return c.value
}

func (c UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateGdgwRoutetableRequestBodyGdgwRoutetableAddRoutesType) UnmarshalJSON(b []byte) error {
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
