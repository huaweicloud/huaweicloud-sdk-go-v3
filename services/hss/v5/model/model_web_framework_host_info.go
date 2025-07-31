package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebFrameworkHostInfo 服务器列表
type WebFrameworkHostInfo struct {

	// agent_id
	AgentId *string `json:"agent_id,omitempty"`

	// 主机id
	HostId *string `json:"host_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器ip
	HostIp *string `json:"host_ip,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`

	// 路径
	Path *string `json:"path,omitempty"`

	// 大小
	FileSize *int32 `json:"file_size,omitempty"`

	// 扫描时间
	RecordTime *int64 `json:"record_time,omitempty"`

	// 绑定的ip列表
	BindIpList *string `json:"bind_ip_list,omitempty"`

	// 软件的类型
	Catalogue *string `json:"catalogue,omitempty"`

	// 连接的ip列表
	ConnectedIpList *string `json:"connected_ip_list,omitempty"`

	// 连接数
	ConnectedNumber *string `json:"connected_number,omitempty"`

	// 压缩的目录
	EmbedderDir *string `json:"embedder_dir,omitempty"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// 文件类型
	FileType *string `json:"file_type,omitempty"`

	// 用户组id
	Gid *int32 `json:"gid,omitempty"`

	// 文件哈希值
	Hash *string `json:"hash,omitempty"`

	// 是否是压缩的文件
	IsEmbedded *int32 `json:"is_embedded,omitempty"`

	// 监听的端口列表
	ListenPortList *string `json:"listen_port_list,omitempty"`

	// 文件权限
	Mode *string `json:"mode,omitempty"`

	// 进程id
	Pid *int32 `json:"pid,omitempty"`

	// 进程路径
	ProcPath *string `json:"proc_path,omitempty"`

	// 用户id
	Uid *string `json:"uid,omitempty"`
}

func (o WebFrameworkHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebFrameworkHostInfo struct{}"
	}

	return strings.Join([]string{"WebFrameworkHostInfo", string(data)}, " ")
}
