package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulHostProcessResponseInfoDataList struct {

	// **参数解释**： 进程id **取值范围**： 字符长度0-256位
	Pid *string `json:"pid,omitempty"`

	// **参数解释**： 绑定ip **取值范围**： 字符长度0-128位
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**： 自启动项的路径 **取值范围**： 字符长度0-256位
	Path *string `json:"path,omitempty"`

	// **参数解释**： 端口/协议 **取值范围**： 字符长度0-32位
	PortProtocol *string `json:"port_protocol,omitempty"`

	// **参数解释**： 版本 **取值范围**： 字符长度0-256位
	Version *string `json:"version,omitempty"`

	// **参数解释**： 依赖包 **取值范围**： 字符长度0-256位
	DependencyPackage *string `json:"dependency_package,omitempty"`

	// **参数解释**： cpu使用率 **取值范围**： 字符长度0-32位
	CpuUseRate *string `json:"cpu_use_rate,omitempty"`
}

func (o VulHostProcessResponseInfoDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulHostProcessResponseInfoDataList struct{}"
	}

	return strings.Join([]string{"VulHostProcessResponseInfoDataList", string(data)}, " ")
}
