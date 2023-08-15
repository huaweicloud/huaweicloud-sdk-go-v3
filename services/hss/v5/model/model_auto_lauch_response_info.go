package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoLauchResponseInfo 中间件信息
type AutoLauchResponseInfo struct {

	// agent_id
	AgentId *string `json:"agent_id,omitempty"`

	// 主机id
	HostId *string `json:"host_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器ip
	HostIp *string `json:"host_ip,omitempty"`

	// 自启动项名称
	Name *string `json:"name,omitempty"`

	// 自启动项类型
	Type *int32 `json:"type,omitempty"`

	// 路径
	Path *string `json:"path,omitempty"`

	// 文件hash
	Hash *string `json:"hash,omitempty"`

	// 运行用户
	RunUser *string `json:"run_user,omitempty"`

	// 最近扫描时间
	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`
}

func (o AutoLauchResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoLauchResponseInfo struct{}"
	}

	return strings.Join([]string{"AutoLauchResponseInfo", string(data)}, " ")
}
