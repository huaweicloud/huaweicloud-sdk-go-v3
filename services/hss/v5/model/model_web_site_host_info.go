package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebSiteHostInfo 服务器列表
type WebSiteHostInfo struct {

	// agent_id
	AgentId *string `json:"agent_id,omitempty"`

	// 主机id
	HostId *string `json:"host_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器ip
	HostIp *string `json:"host_ip,omitempty"`

	// 域名
	Domain *string `json:"domain,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 路径
	Path *string `json:"path,omitempty"`

	// 端口
	Port *int32 `json:"port,omitempty"`

	// 绑定地址
	BindAddr *string `json:"bind_addr,omitempty"`

	// url路径
	UrlPath *string `json:"url_path,omitempty"`

	// 用户id
	Uid *int32 `json:"uid,omitempty"`

	// 用户组id
	Gid *int32 `json:"gid,omitempty"`

	// 文件权限
	Mode *string `json:"mode,omitempty"`

	// 进程id
	Pid *int32 `json:"pid,omitempty"`

	// 进程路径
	ProcPath *string `json:"proc_path,omitempty"`

	// 是否是https
	IsHttps *bool `json:"is_https,omitempty"`

	// SSL证书颁发者
	CertIssuer *string `json:"cert_issuer,omitempty"`

	// SSL证书使用者
	CertUser *string `json:"cert_user,omitempty"`

	// SSL证书颁发时间
	CertIssueTime *string `json:"cert_issue_time,omitempty"`

	// SSL证书截止时间
	CertExpiredTime *string `json:"cert_expired_time,omitempty"`

	// 扫描时间
	RecordTime *int64 `json:"record_time,omitempty"`

	// 容器id
	ContainerId *string `json:"container_id,omitempty"`

	// 容器名称
	ContainerName *string `json:"container_name,omitempty"`
}

func (o WebSiteHostInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebSiteHostInfo struct{}"
	}

	return strings.Join([]string{"WebSiteHostInfo", string(data)}, " ")
}
