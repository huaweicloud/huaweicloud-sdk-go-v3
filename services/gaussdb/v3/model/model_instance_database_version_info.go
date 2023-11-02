package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceDatabaseVersionInfo 数据库版本信息。
type InstanceDatabaseVersionInfo struct {

	// 当前数据库版本。
	CurrentVersion *string `json:"current_version,omitempty"`

	// 当前数据库内核版本。
	CurrentKernelVersion *string `json:"current_kernel_version,omitempty"`

	// 最新数据库版本。
	LatestVersion *string `json:"latest_version,omitempty"`

	// 最新数据库内核版本。
	LatestKernelVersion *string `json:"latest_kernel_version,omitempty"`
}

func (o InstanceDatabaseVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDatabaseVersionInfo struct{}"
	}

	return strings.Join([]string{"InstanceDatabaseVersionInfo", string(data)}, " ")
}
