package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileRedirectionOptionsLinuxFileSizeSupportedOptions Linux支持设置文件大小控制项。
type FileRedirectionOptionsLinuxFileSizeSupportedOptions struct {

	// Linux设置文件大小阈值（Byte）。取值范围为[0-4096]。默认：100。
	LinuxFileSizeSupportedThreshold *int32 `json:"linux_file_size_supported_threshold,omitempty"`
}

func (o FileRedirectionOptionsLinuxFileSizeSupportedOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileRedirectionOptionsLinuxFileSizeSupportedOptions struct{}"
	}

	return strings.Join([]string{"FileRedirectionOptionsLinuxFileSizeSupportedOptions", string(data)}, " ")
}
