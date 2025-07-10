package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPersonalizedDataMgmt 个性化数据管理。
type PoliciesPersonalizedDataMgmt struct {

	// 个性化数据管理路径。
	PersonalizedDataMgmtPath *string `json:"personalized_data_mgmt_path,omitempty"`

	// 用户数据漫游。
	UserDataRoamingEnable *bool `json:"user_data_roaming_enable,omitempty"`

	UserDataRoamingOptions *PoliciesPersonalizedDataMgmtUserDataRoamingOptions `json:"user_data_roaming_options,omitempty"`

	// 启用用户文件夹重定向。
	UserFolderRedirectionEnable *bool `json:"user_folder_redirection_enable,omitempty"`

	UserFolderRedirectionOptions *PoliciesPersonalizedDataMgmtUserFolderRedirectionOptions `json:"user_folder_redirection_options,omitempty"`

	// 启用用户文件夹重定向。
	LogoffDeleteUserConfiguration *bool `json:"logoff_delete_user_configuration,omitempty"`

	// 启用用户文件夹重定向。
	NetworkDriveMappingEnable *bool `json:"network_drive_mapping_enable,omitempty"`

	NetworkDriveMappingOptions *PoliciesPersonalizedDataMgmtNetworkDriveMappingOptions `json:"network_drive_mapping_options,omitempty"`
}

func (o PoliciesPersonalizedDataMgmt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPersonalizedDataMgmt struct{}"
	}

	return strings.Join([]string{"PoliciesPersonalizedDataMgmt", string(data)}, " ")
}
