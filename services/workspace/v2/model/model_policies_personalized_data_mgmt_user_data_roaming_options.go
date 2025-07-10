package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPersonalizedDataMgmtUserDataRoamingOptions 用户数据漫游选项。
type PoliciesPersonalizedDataMgmtUserDataRoamingOptions struct {

	// 配置文件流式处理启用。
	ProfileStreamingEnable *bool `json:"profile_streaming_enable,omitempty"`

	// 漫游文件本地路径。
	RoamedFilesLocalPath *string `json:"roamed_files_local_path,omitempty"`

	// 排除文件夹路径。
	ExcludeFoldersPath *string `json:"exclude_folders_path,omitempty"`

	// 排除文件夹路径。
	RoamingRegistryMethod *string `json:"roaming_registry_method,omitempty"`

	// 漫游注册路径。
	RoamingRegistryPath *string `json:"roaming_registry_path,omitempty"`
}

func (o PoliciesPersonalizedDataMgmtUserDataRoamingOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPersonalizedDataMgmtUserDataRoamingOptions struct{}"
	}

	return strings.Join([]string{"PoliciesPersonalizedDataMgmtUserDataRoamingOptions", string(data)}, " ")
}
