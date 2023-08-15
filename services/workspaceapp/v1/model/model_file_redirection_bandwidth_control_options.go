package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FileRedirectionBandwidthControlOptions struct {

	// 文件重定向带宽控制量（Kbps）。取值范围为[500-20000]。默认：10000。
	FileRedirectionBandwidthControlValue *int32 `json:"file_redirection_bandwidth_control_value,omitempty"`
}

func (o FileRedirectionBandwidthControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileRedirectionBandwidthControlOptions struct{}"
	}

	return strings.Join([]string{"FileRedirectionBandwidthControlOptions", string(data)}, " ")
}
