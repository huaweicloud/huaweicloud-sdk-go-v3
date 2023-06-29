package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FileRedirectionBandwidthPercentageOptions struct {

	// 文件重定向带宽百分比控制量（%）。取值范围为[0-100]。默认：30。
	FileRedirectionBandwidthPercentageValue *int32 `json:"file_redirection_bandwidth_percentage_value,omitempty"`
}

func (o FileRedirectionBandwidthPercentageOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileRedirectionBandwidthPercentageOptions struct{}"
	}

	return strings.Join([]string{"FileRedirectionBandwidthPercentageOptions", string(data)}, " ")
}
