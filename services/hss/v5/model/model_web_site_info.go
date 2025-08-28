package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WebSiteInfo **参数解释**: Web站点信息
type WebSiteInfo struct {

	// **参数解释**: 域名 **取值范围**: 字符长度1-256
	Domain *string `json:"domain,omitempty"`

	// **参数解释**: 软件名称 **取值范围**: 字符长度1-256
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: web目录路径 **取值范围**: 字符长度1-512
	Path *string `json:"path,omitempty"`

	// **参数解释**: web站点对外端口 **取值范围**: 最小值0，最大值2147483647
	Port *int32 `json:"port,omitempty"`

	// **参数解释**: web站点绑定IP **取值范围**: 字符长度1-512
	BindAddr *string `json:"bind_addr,omitempty"`

	// **参数解释**: web站点url路径 **取值范围**: 字符长度1-128
	UrlPath *string `json:"url_path,omitempty"`

	// **参数解释**: web站点进程uid **取值范围**: 最小值0，最大值2147483647
	Uid *int32 `json:"uid,omitempty"`

	// **参数解释**: web站点进程gid **取值范围**: 最小值0，最大值2147483647
	Gid *int32 `json:"gid,omitempty"`

	// **参数解释**: web站点文件权限 **取值范围**: 字符长度1-32
	Mode *string `json:"mode,omitempty"`

	// **参数解释**: web站点进程pid **取值范围**: 最小值0，最大值2147483647
	Pid *int32 `json:"pid,omitempty"`

	// **参数解释**: web站点进程路径 **取值范围**: 字符长度1-1024
	ProcPath *string `json:"proc_path,omitempty"`

	// **参数解释**: web站点是否为https **取值范围**: -true：是。 -false：否。
	IsHttps *bool `json:"is_https,omitempty"`

	// **参数解释**: web站点SSL证书颁发者 **取值范围**: 字符长度0-256
	CertIssuer *string `json:"cert_issuer,omitempty"`

	// **参数解释**: web站点SSL证书使用者 **取值范围**: 字符长度0-256
	CertUser *int32 `json:"cert_user,omitempty"`

	// **参数解释**: web站点SSL证书颁发时间 **取值范围**: 字符长度0-32
	CertIssueTime *string `json:"cert_issue_time,omitempty"`

	// **参数解释**: web站点SSL证书截止时间 **取值范围**: 字符长度0-32
	CertExpiredTime *string `json:"cert_expired_time,omitempty"`

	// **参数解释**: web框架扫描时间 **取值范围**: 最小值0，最大值2^63-1
	RecordTime *int64 `json:"record_time,omitempty"`

	// **参数解释**: 容器id **取值范围**: 字符长度1-128
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**: 容器名称 **取值范围**: 字符长度1-256
	ContainerName *string `json:"container_name,omitempty"`
}

func (o WebSiteInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebSiteInfo struct{}"
	}

	return strings.Join([]string{"WebSiteInfo", string(data)}, " ")
}
