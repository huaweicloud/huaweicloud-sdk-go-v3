package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// MicroServiceInfo 微服务的响应对象
type MicroServiceInfo struct {

	// 微服务编号
	Id *string `json:"id,omitempty"`

	// 实例编号
	InstanceId *string `json:"instance_id,omitempty"`

	// 微服务类型： - CSE：CSE微服务注册中心 - CCE：CCE云容器引擎（工作负载） - CCE_SERVICE: CCE云容器引擎（Service）
	ServiceType *MicroServiceInfoServiceType `json:"service_type,omitempty"`

	CseInfo *MicroServiceInfoCse `json:"cse_info,omitempty"`

	CceInfo *MicroServiceInfoCce `json:"cce_info,omitempty"`

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
	CSE         MicroServiceInfoServiceType
	CCE         MicroServiceInfoServiceType
	CCE_SERVICE MicroServiceInfoServiceType
}

func GetMicroServiceInfoServiceTypeEnum() MicroServiceInfoServiceTypeEnum {
	return MicroServiceInfoServiceTypeEnum{
		CSE: MicroServiceInfoServiceType{
			value: "CSE",
		},
		CCE: MicroServiceInfoServiceType{
			value: "CCE",
		},
		CCE_SERVICE: MicroServiceInfoServiceType{
			value: "CCE_SERVICE",
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
