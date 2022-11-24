package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MicroServiceCreate struct {

	// 微服务类型： - CSE：CSE微服务注册中心 - CCE：CCE云容器引擎
	ServiceType *MicroServiceCreateServiceType `json:"service_type,omitempty"`

	CseInfo *MicroServiceInfoCseBase `json:"cse_info,omitempty"`

	CceInfo *MicroServiceInfoCceBase `json:"cce_info,omitempty"`
}

func (o MicroServiceCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroServiceCreate struct{}"
	}

	return strings.Join([]string{"MicroServiceCreate", string(data)}, " ")
}

type MicroServiceCreateServiceType struct {
	value string
}

type MicroServiceCreateServiceTypeEnum struct {
	CSE MicroServiceCreateServiceType
	CCE MicroServiceCreateServiceType
}

func GetMicroServiceCreateServiceTypeEnum() MicroServiceCreateServiceTypeEnum {
	return MicroServiceCreateServiceTypeEnum{
		CSE: MicroServiceCreateServiceType{
			value: "CSE",
		},
		CCE: MicroServiceCreateServiceType{
			value: "CCE",
		},
	}
}

func (c MicroServiceCreateServiceType) Value() string {
	return c.value
}

func (c MicroServiceCreateServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MicroServiceCreateServiceType) UnmarshalJSON(b []byte) error {
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
