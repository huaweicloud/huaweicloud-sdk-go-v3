package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 主机信息
type DeploymentHostRequest struct {
	// 主机名称

	HostName string `json:"host_name"`
	// IP，请输入弹性ip格式：161.17.101.12

	Ip string `json:"ip"`
	// ssh端口，如：22

	Port int32 `json:"port"`
	// 是否为代理机

	AsProxy bool `json:"as_proxy"`
	// 代理机id

	ProxyHostId *string `json:"proxy_host_id,omitempty"`

	Authorization *DeploymentHostAuthorizationBody `json:"authorization"`
	// 免费启用应用运维服务（AOM），提供指标监控、日志查询、告警功能（自动安装数据采集器 ICAgent，仅支持华为云linux主机）

	InstallIcagent *bool `json:"install_icagent,omitempty"`
	// 是否将当前主机的密码信息同步到同一项目下其他主机组中具有相同IP、用户名、端口的主机。

	Sync *bool `json:"sync,omitempty"`
}

func (o DeploymentHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeploymentHostRequest struct{}"
	}

	return strings.Join([]string{"DeploymentHostRequest", string(data)}, " ")
}
