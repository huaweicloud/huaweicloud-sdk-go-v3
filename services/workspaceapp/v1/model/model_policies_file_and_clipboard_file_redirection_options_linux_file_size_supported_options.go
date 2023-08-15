package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesFileAndClipboardFileRedirectionOptionsLinuxFileSizeSupportedOptions Linux支持设置文件大小控制项。
type PoliciesFileAndClipboardFileRedirectionOptionsLinuxFileSizeSupportedOptions struct {

	// Linux设置文件大小阈值（Byte）。取值范围为[0-4096]。默认：100。
	LinuxFileSizeSupportedThreshold *int32 `json:"linux_file_size_supported_threshold,omitempty"`
}

func (o PoliciesFileAndClipboardFileRedirectionOptionsLinuxFileSizeSupportedOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesFileAndClipboardFileRedirectionOptionsLinuxFileSizeSupportedOptions struct{}"
	}

	return strings.Join([]string{"PoliciesFileAndClipboardFileRedirectionOptionsLinuxFileSizeSupportedOptions", string(data)}, " ")
}
