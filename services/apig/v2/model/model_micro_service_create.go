package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MicroServiceCreate 微服务详情。
type MicroServiceCreate struct {

	// 微服务类型： - CSE：CSE微服务注册中心 - CCE：CCE云容器引擎（工作负载） - CCE_SERVICE: CCE云容器引擎（Service）
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
	CSE         MicroServiceCreateServiceType
	CCE         MicroServiceCreateServiceType
	CCE_SERVICE MicroServiceCreateServiceType
}

func GetMicroServiceCreateServiceTypeEnum() MicroServiceCreateServiceTypeEnum {
	return MicroServiceCreateServiceTypeEnum{
		CSE: MicroServiceCreateServiceType{
			value: "CSE",
		},
		CCE: MicroServiceCreateServiceType{
			value: "CCE",
		},
		CCE_SERVICE: MicroServiceCreateServiceType{
			value: "CCE_SERVICE",
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
