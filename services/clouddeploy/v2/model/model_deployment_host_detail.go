package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 主机信息详情
type DeploymentHostDetail struct {
	// 主机组id

	GroupId string `json:"group_id"`
	// 主机名称

	HostName string `json:"host_name"`
	// IP，请输入弹性ip格式：161.17.101.12

	Ip string `json:"ip"`
	// ssh端口，如：22

	Port int32 `json:"port"`
	// 操作系统：windows|linux，需要和主机组保持一致

	Os DeploymentHostDetailOs `json:"os"`
	// 是否为代理机

	AsProxy bool `json:"as_proxy"`
	// 代理机id

	ProxyHostId *string `json:"proxy_host_id,omitempty"`

	Authorization *DeploymentHostAuthorizationBody `json:"authorization"`
	// 免费启用应用运维服务（AOM），提供指标监控、日志查询、告警功能（自动安装数据采集器 ICAgent，仅支持华为云linux主机）

	InstallIcagent *bool `json:"install_icagent,omitempty"`
	// 主机ID

	HostId *string `json:"host_id,omitempty"`

	ProxyHost *DeploymentHostDetail `json:"proxy_host,omitempty"`
	// 主机组名

	GroupName *string `json:"group_name,omitempty"`
	// devcloud项目id

	ProjectId *string `json:"project_id,omitempty"`
	// devcloud项目名称

	ProjectName *string `json:"project_name,omitempty"`

	Permission *PermissionHostDetail `json:"permission,omitempty"`
}

func (o DeploymentHostDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentHostDetail struct{}"
	}

	return strings.Join([]string{"DeploymentHostDetail", string(data)}, " ")
}

type DeploymentHostDetailOs struct {
	value string
}

type DeploymentHostDetailOsEnum struct {
	WINDOWS DeploymentHostDetailOs
	LINUX   DeploymentHostDetailOs
}

func GetDeploymentHostDetailOsEnum() DeploymentHostDetailOsEnum {
	return DeploymentHostDetailOsEnum{
		WINDOWS: DeploymentHostDetailOs{
			value: "windows",
		},
		LINUX: DeploymentHostDetailOs{
			value: "linux",
		},
	}
}

func (c DeploymentHostDetailOs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeploymentHostDetailOs) UnmarshalJSON(b []byte) error {
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
