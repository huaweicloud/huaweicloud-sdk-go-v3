package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BackendConfig struct {

	// 后端请求类型
	Type *BackendConfigType `json:"type,omitempty"`

	// 后端请求协议类型
	Protocol *BackendConfigProtocol `json:"protocol,omitempty"`

	// 后端host
	Host *string `json:"host,omitempty"`

	// 后端超时时间
	Timeout *int32 `json:"timeout,omitempty"`

	// 后端请求Path
	Path *string `json:"path,omitempty"`

	// API后端参数
	BackendParas *[]BackendRequestPara `json:"backend_paras,omitempty"`

	// 后端常量参数
	ConstantParas *[]BackendConstant `json:"constant_paras,omitempty"`
}

func (o BackendConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackendConfig struct{}"
	}

	return strings.Join([]string{"BackendConfig", string(data)}, " ")
}

type BackendConfigType struct {
	value string
}

type BackendConfigTypeEnum struct {
	REQUEST_TYPE_POST BackendConfigType
	REQUEST_TYPE_GET  BackendConfigType
}

func GetBackendConfigTypeEnum() BackendConfigTypeEnum {
	return BackendConfigTypeEnum{
		REQUEST_TYPE_POST: BackendConfigType{
			value: "REQUEST_TYPE_POST",
		},
		REQUEST_TYPE_GET: BackendConfigType{
			value: "REQUEST_TYPE_GET",
		},
	}
}

func (c BackendConfigType) Value() string {
	return c.value
}

func (c BackendConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendConfigType) UnmarshalJSON(b []byte) error {
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

type BackendConfigProtocol struct {
	value string
}

type BackendConfigProtocolEnum struct {
	PROTOCOL_TYPE_HTTP  BackendConfigProtocol
	PROTOCOL_TYPE_HTTPS BackendConfigProtocol
}

func GetBackendConfigProtocolEnum() BackendConfigProtocolEnum {
	return BackendConfigProtocolEnum{
		PROTOCOL_TYPE_HTTP: BackendConfigProtocol{
			value: "PROTOCOL_TYPE_HTTP",
		},
		PROTOCOL_TYPE_HTTPS: BackendConfigProtocol{
			value: "PROTOCOL_TYPE_HTTPS",
		},
	}
}

func (c BackendConfigProtocol) Value() string {
	return c.value
}

func (c BackendConfigProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackendConfigProtocol) UnmarshalJSON(b []byte) error {
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
