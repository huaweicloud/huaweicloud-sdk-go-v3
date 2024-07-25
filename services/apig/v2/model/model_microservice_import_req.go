package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MicroserviceImportReq 导入微服务的请求对象
type MicroserviceImportReq struct {
	GroupInfo *MicroserviceGroup `json:"group_info"`

	// 微服务中心类型。 - CSE：CSE微服务注册中心 - CCE: CCE云容器引擎（工作负载） - CCE_SERVICE: CCE云容器引擎（Service）
	ServiceType MicroserviceImportReqServiceType `json:"service_type"`

	// API网关访问微服务的请求协议 - HTTP - HTTPS
	Protocol *MicroserviceImportReqProtocol `json:"protocol,omitempty"`

	// 导入的api列表
	Apis []MicroserviceApiCreate `json:"apis"`

	// APIG请求后端服务的超时时间。最大超时时间可通过实例特性backend_timeout配置修改，可修改的上限为600000，默认5000  单位：毫秒。
	BackendTimeout *int32 `json:"backend_timeout,omitempty"`

	// API的认证方式，默认无认证。 - NONE：无认证 - APP：APP认证 - IAM：IAM认证
	AuthType *MicroserviceImportReqAuthType `json:"auth_type,omitempty"`

	// 是否支持跨域，默认不支持 - true：支持 - false：不支持
	Cors *bool `json:"cors,omitempty"`

	CseInfo *MicroServiceInfoCseCreate `json:"cse_info,omitempty"`

	CceInfo *MicroServiceInfoCceCreate `json:"cce_info,omitempty"`
}

func (o MicroserviceImportReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MicroserviceImportReq struct{}"
	}

	return strings.Join([]string{"MicroserviceImportReq", string(data)}, " ")
}

type MicroserviceImportReqServiceType struct {
	value string
}

type MicroserviceImportReqServiceTypeEnum struct {
	CSE         MicroserviceImportReqServiceType
	CCE         MicroserviceImportReqServiceType
	CCE_SERVICE MicroserviceImportReqServiceType
}

func GetMicroserviceImportReqServiceTypeEnum() MicroserviceImportReqServiceTypeEnum {
	return MicroserviceImportReqServiceTypeEnum{
		CSE: MicroserviceImportReqServiceType{
			value: "CSE",
		},
		CCE: MicroserviceImportReqServiceType{
			value: "CCE",
		},
		CCE_SERVICE: MicroserviceImportReqServiceType{
			value: "CCE_SERVICE",
		},
	}
}

func (c MicroserviceImportReqServiceType) Value() string {
	return c.value
}

func (c MicroserviceImportReqServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MicroserviceImportReqServiceType) UnmarshalJSON(b []byte) error {
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

type MicroserviceImportReqProtocol struct {
	value string
}

type MicroserviceImportReqProtocolEnum struct {
	HTTP  MicroserviceImportReqProtocol
	HTTPS MicroserviceImportReqProtocol
}

func GetMicroserviceImportReqProtocolEnum() MicroserviceImportReqProtocolEnum {
	return MicroserviceImportReqProtocolEnum{
		HTTP: MicroserviceImportReqProtocol{
			value: "HTTP",
		},
		HTTPS: MicroserviceImportReqProtocol{
			value: "HTTPS",
		},
	}
}

func (c MicroserviceImportReqProtocol) Value() string {
	return c.value
}

func (c MicroserviceImportReqProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MicroserviceImportReqProtocol) UnmarshalJSON(b []byte) error {
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

type MicroserviceImportReqAuthType struct {
	value string
}

type MicroserviceImportReqAuthTypeEnum struct {
	NONE MicroserviceImportReqAuthType
	APP  MicroserviceImportReqAuthType
	IAM  MicroserviceImportReqAuthType
}

func GetMicroserviceImportReqAuthTypeEnum() MicroserviceImportReqAuthTypeEnum {
	return MicroserviceImportReqAuthTypeEnum{
		NONE: MicroserviceImportReqAuthType{
			value: "NONE",
		},
		APP: MicroserviceImportReqAuthType{
			value: "APP",
		},
		IAM: MicroserviceImportReqAuthType{
			value: "IAM",
		},
	}
}

func (c MicroserviceImportReqAuthType) Value() string {
	return c.value
}

func (c MicroserviceImportReqAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MicroserviceImportReqAuthType) UnmarshalJSON(b []byte) error {
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
