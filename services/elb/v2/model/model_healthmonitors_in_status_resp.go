package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HealthmonitorsInStatusResp 健康检查对象，用于状态树中
type HealthmonitorsInStatusResp struct {

	// 健康检查ID
	Id string `json:"id"`

	// 健康检查名称
	Name string `json:"name"`

	// 健康检查类型
	Type HealthmonitorsInStatusRespType `json:"type"`

	// 健康检查的配置状态；该字段为预留字段，暂未启用。默认为ACTIVE。
	ProvisioningStatus string `json:"provisioning_status"`
}

func (o HealthmonitorsInStatusResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthmonitorsInStatusResp struct{}"
	}

	return strings.Join([]string{"HealthmonitorsInStatusResp", string(data)}, " ")
}

type HealthmonitorsInStatusRespType struct {
	value string
}

type HealthmonitorsInStatusRespTypeEnum struct {
	UDP_CONNECT HealthmonitorsInStatusRespType
	TCP         HealthmonitorsInStatusRespType
	HTTP        HealthmonitorsInStatusRespType
}

func GetHealthmonitorsInStatusRespTypeEnum() HealthmonitorsInStatusRespTypeEnum {
	return HealthmonitorsInStatusRespTypeEnum{
		UDP_CONNECT: HealthmonitorsInStatusRespType{
			value: "UDP_CONNECT",
		},
		TCP: HealthmonitorsInStatusRespType{
			value: "TCP",
		},
		HTTP: HealthmonitorsInStatusRespType{
			value: "HTTP",
		},
	}
}

func (c HealthmonitorsInStatusRespType) Value() string {
	return c.value
}

func (c HealthmonitorsInStatusRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HealthmonitorsInStatusRespType) UnmarshalJSON(b []byte) error {
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
