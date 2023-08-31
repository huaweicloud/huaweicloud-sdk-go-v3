package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateHostRequestBody 主机信息body体
type CreateHostRequestBody struct {

	// 主机名称
	HostName string `json:"host_name"`

	// IP，请输入弹性ip格式：161.17.101.12
	Ip string `json:"ip"`

	// ssh端口，如：22
	Port int32 `json:"port"`

	// 操作系统：windows|linux，需要和主机集群保持一致
	Os CreateHostRequestBodyOs `json:"os"`

	// 是否为代理机
	AsProxy bool `json:"as_proxy"`

	// 代理机id
	ProxyHostId *string `json:"proxy_host_id,omitempty"`

	Authorization *HostAuthorizationBody `json:"authorization"`

	// 免费启用应用运维服务（AOM），提供指标监控、日志查询、告警功能（自动安装数据采集器 ICAgent，仅支持华为云linux主机）
	InstallIcagent *bool `json:"install_icagent,omitempty"`
}

func (o CreateHostRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostRequestBody struct{}"
	}

	return strings.Join([]string{"CreateHostRequestBody", string(data)}, " ")
}

type CreateHostRequestBodyOs struct {
	value string
}

type CreateHostRequestBodyOsEnum struct {
	WINDOWS CreateHostRequestBodyOs
	LINUX   CreateHostRequestBodyOs
}

func GetCreateHostRequestBodyOsEnum() CreateHostRequestBodyOsEnum {
	return CreateHostRequestBodyOsEnum{
		WINDOWS: CreateHostRequestBodyOs{
			value: "windows",
		},
		LINUX: CreateHostRequestBodyOs{
			value: "linux",
		},
	}
}

func (c CreateHostRequestBodyOs) Value() string {
	return c.value
}

func (c CreateHostRequestBodyOs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHostRequestBodyOs) UnmarshalJSON(b []byte) error {
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
