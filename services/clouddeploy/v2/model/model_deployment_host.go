package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 主机信息body体
type DeploymentHost struct {
	// 主机名称

	HostName string `json:"host_name"`
	// IP，请输入弹性ip格式：161.17.101.12

	Ip string `json:"ip"`
	// ssh端口，如：22

	Port int32 `json:"port"`
	// 操作系统：windows|linux，需要和主机组保持一致

	Os DeploymentHostOs `json:"os"`
	// 是否为代理机

	AsProxy bool `json:"as_proxy"`
	// 代理机id

	ProxyHostId *string `json:"proxy_host_id,omitempty"`

	Authorization *DeploymentHostAuthorizationBody `json:"authorization"`
	// 免费启用应用运维服务（AOM），提供指标监控、日志查询、告警功能（自动安装数据采集器 ICAgent，仅支持华为云linux主机）

	InstallIcagent *bool `json:"install_icagent,omitempty"`
}

func (o DeploymentHost) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentHost struct{}"
	}

	return strings.Join([]string{"DeploymentHost", string(data)}, " ")
}

type DeploymentHostOs struct {
	value string
}

type DeploymentHostOsEnum struct {
	WINDOWS DeploymentHostOs
	LINUX   DeploymentHostOs
}

func GetDeploymentHostOsEnum() DeploymentHostOsEnum {
	return DeploymentHostOsEnum{
		WINDOWS: DeploymentHostOs{
			value: "windows",
		},
		LINUX: DeploymentHostOs{
			value: "linux",
		},
	}
}

func (c DeploymentHostOs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeploymentHostOs) UnmarshalJSON(b []byte) error {
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
