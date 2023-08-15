package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoLaunchChangeResponseInfo 自启动项变动历史信息
type AutoLaunchChangeResponseInfo struct {

	// agent_id
	AgentId *string `json:"agent_id,omitempty"`

	// the type of change   - add ：新建   - delete ：删除   - modify ：修改
	VariationType *string `json:"variation_type,omitempty"`

	// 自启动项类型
	Type *int32 `json:"type,omitempty"`

	// host_id
	HostId *string `json:"host_id,omitempty"`

	// 弹性服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 主机IP
	HostIp *string `json:"host_ip,omitempty"`

	// 路径
	Path *string `json:"path,omitempty"`

	// 文件hash
	Hash *string `json:"hash,omitempty"`

	// 运行用户
	RunUser *string `json:"run_user,omitempty"`

	// 自启动项名称
	Name *string `json:"name,omitempty"`

	// 最近更新时间
	RecentScanTime *int64 `json:"recent_scan_time,omitempty"`
}

func (o AutoLaunchChangeResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoLaunchChangeResponseInfo struct{}"
	}

	return strings.Join([]string{"AutoLaunchChangeResponseInfo", string(data)}, " ")
}
