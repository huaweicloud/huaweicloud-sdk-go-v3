package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// MicroServiceInfo 微服务的响应对象
type MicroServiceInfo struct {

	// 微服务编号
	Id *string `json:"id,omitempty"`

	// 微服务类型： - CSE：CSE微服务注册中心
	ServiceType *MicroServiceInfoServiceType `json:"service_type,omitempty"`

	CseInfo *MicroServiceInfoCse `json:"cse_info,omitempty"`

	// 微服务更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 微服务创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
}

func (o MicroServiceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroServiceInfo struct{}"
	}

	return strings.Join([]string{"MicroServiceInfo", string(data)}, " ")
}

type MicroServiceInfoServiceType struct {
	value string
}

type MicroServiceInfoServiceTypeEnum struct {
	CSE MicroServiceInfoServiceType
}

func GetMicroServiceInfoServiceTypeEnum() MicroServiceInfoServiceTypeEnum {
	return MicroServiceInfoServiceTypeEnum{
		CSE: MicroServiceInfoServiceType{
			value: "CSE",
		},
	}
}

func (c MicroServiceInfoServiceType) Value() string {
	return c.value
}

func (c MicroServiceInfoServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MicroServiceInfoServiceType) UnmarshalJSON(b []byte) error {
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
