package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TenantVersionInfo 服务端版本信息
type TenantVersionInfo struct {

	// 版本ID
	Id *string `json:"id,omitempty"`

	// 版本号
	Version *string `json:"version,omitempty"`

	// 版本类型：0-服务端 1-客户端
	VersionType *int32 `json:"version_type,omitempty"`

	// 操作系统类型：0-windows 1-android 2-mac 3-linux_UOS 4-linux_ubuntu 5-linux_Kylin 6-linux 7-linux_ubuntu_soft 8-linux_kylin_v10
	OsType *int32 `json:"os_type,omitempty"`

	// 更新说明
	ReleaseNote *string `json:"release_note,omitempty"`

	// 租户自定义更新说明
	CustomReleaseNote *string `json:"custom_release_note,omitempty"`

	// 版本下载地址
	VersionDownloadUrl *string `json:"version_download_url,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 版本状态：PREVIEW-预览 RELEASED-已发布 OFFLINE-已下线 OBSOLETE-已废弃
	VersionStatus *string `json:"version_status,omitempty"`

	// 发布时间
	PublishTime *string `json:"publish_time,omitempty"`

	// 停止服务时间
	StopTime *string `json:"stop_time,omitempty"`
}

func (o TenantVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantVersionInfo struct{}"
	}

	return strings.Join([]string{"TenantVersionInfo", string(data)}, " ")
}
