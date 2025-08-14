package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoLauchResponseInfo 中间件信息
type AutoLauchResponseInfo struct {

	// **参数解释**: Agent ID **取值范围**: 字符长度1-64位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 主机ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 主机IP **取值范围**: 字符长度1-128位
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 自启动项名称 **取值范围**: 字符长度1-256位
	Name *string `json:"name,omitempty"`

	// **参数解释**: 自启动项类型 **取值范围**: - 0：自启动服务 - 1：定时任务 - 2：预加载动态库 - 3：Run注册表键 - 4：开机启动文件夹
	Type *int32 `json:"type,omitempty"`

	// **参数解释**: 自启动项的路径 **取值范围**: 字符长度1-256位
	Path *string `json:"path,omitempty"`

	// **参数解释**: 采用sha256算法生成的文件hash值 **取值范围**: 字符长度1-128位
	Hash *string `json:"hash,omitempty"`

	// **参数解释**: 运行用户 **取值范围**: 字符长度1-128位
	RunUser *string `json:"run_user,omitempty"`

	// **参数解释**: 最近扫描时间 **取值范围**: 最小值0，最大值9223372036854775807
	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`
}

func (o AutoLauchResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoLauchResponseInfo struct{}"
	}

	return strings.Join([]string{"AutoLauchResponseInfo", string(data)}, " ")
}
