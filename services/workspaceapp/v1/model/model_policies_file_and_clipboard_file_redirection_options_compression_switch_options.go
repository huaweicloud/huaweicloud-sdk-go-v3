package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesFileAndClipboardFileRedirectionOptionsCompressionSwitchOptions 压缩开关控制项。
type PoliciesFileAndClipboardFileRedirectionOptionsCompressionSwitchOptions struct {

	// 压缩阈值（Byte）。取值范围为[0-10240]。默认：512。
	CompressionThreshold *int32 `json:"compression_threshold,omitempty"`

	// 最小压缩率。取值范围为[0-1000]。默认：900。
	MinimumCompressionRate *int32 `json:"minimum_compression_rate,omitempty"`
}

func (o PoliciesFileAndClipboardFileRedirectionOptionsCompressionSwitchOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesFileAndClipboardFileRedirectionOptionsCompressionSwitchOptions struct{}"
	}

	return strings.Join([]string{"PoliciesFileAndClipboardFileRedirectionOptionsCompressionSwitchOptions", string(data)}, " ")
}
