package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantVersionConfigRequest Request Object
type ListTenantVersionConfigRequest struct {

	// 版本类型：0-服务端 1-客户端（必传）
	VersionType int32 `json:"version_type"`

	// 版本号（支持模糊查询）
	Version *string `json:"version,omitempty"`

	// 操作系统类型：0-windows 1-android 2-mac 3-linux_UOS 4-linux_ubuntu 5-linux_Kylin 6-linux
	OsType *int32 `json:"os_type,omitempty"`

	// 版本状态：PREVIEW-预览 RELEASED-已发布 OFFLINE-已下线 OBSOLETE-已废弃
	VersionStatus *string `json:"version_status,omitempty"`

	// 发布时间开始（格式：yyyy-MM-dd HH:mm:ss）
	PublishTimeBegin *string `json:"publish_time_begin,omitempty"`

	// 发布时间结束（格式：yyyy-MM-dd HH:mm:ss）
	PublishTimeEnd *string `json:"publish_time_end,omitempty"`

	// 版本说明（支持模糊查询，会同时搜索SRE配置的版本说明和租户自定义的版本说明）
	ReleaseNote *string `json:"release_note,omitempty"`

	// 偏移量，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量，默认10，最大20000
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTenantVersionConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantVersionConfigRequest struct{}"
	}

	return strings.Join([]string{"ListTenantVersionConfigRequest", string(data)}, " ")
}
