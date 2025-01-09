package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPersonalizedDataMgmtNetworkDriveMappingOptions 网络驱动器映射选项。
type PoliciesPersonalizedDataMgmtNetworkDriveMappingOptions struct {

	// 网络驱动器映射路径
	NetworkDriveMappingPath *string `json:"network_drive_mapping_path,omitempty"`

	// 排除文件夹路径
	DriveLetter *string `json:"drive_letter,omitempty"`
}

func (o PoliciesPersonalizedDataMgmtNetworkDriveMappingOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPersonalizedDataMgmtNetworkDriveMappingOptions struct{}"
	}

	return strings.Join([]string{"PoliciesPersonalizedDataMgmtNetworkDriveMappingOptions", string(data)}, " ")
}
