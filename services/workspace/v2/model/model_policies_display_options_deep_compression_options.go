package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesDisplayOptionsDeepCompressionOptions 深度压缩控制选项。
type PoliciesDisplayOptionsDeepCompressionOptions struct {

	// 深度压缩级别。取值为： 压缩级别0：Compression grade 0 压缩级别1：Compression grade 1 压缩级别2：Compression grade 2 压缩级别3：Compression grade 3 压缩级别4：Compression grade 4 压缩级别5：Compression grade 5 压缩级别6：Compression grade 6 压缩级别7：Compression grade 7 压缩级别8：Compression grade 8 压缩级别9：Compression grade 9
	DeepCompressionLevel *string `json:"deep_compression_level,omitempty"`
}

func (o PoliciesDisplayOptionsDeepCompressionOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesDisplayOptionsDeepCompressionOptions struct{}"
	}

	return strings.Join([]string{"PoliciesDisplayOptionsDeepCompressionOptions", string(data)}, " ")
}
