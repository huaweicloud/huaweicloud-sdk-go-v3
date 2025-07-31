package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddBaselineWhiteListRequestBodyRuleList 待添加的白名单的检查项信息
type AddBaselineWhiteListRequestBodyRuleList struct {

	// 基线检查的检查项标识
	IndexVersion *string `json:"index_version,omitempty"`

	// 基线检查的基线名称
	CheckType *string `json:"check_type,omitempty"`

	// 标准类型，包含如下:   - cn_standard : 等保合规标准   - hw_standard : 云安全实践标准   - cis_standard : 通用安全标准
	Standard *string `json:"standard,omitempty"`
}

func (o AddBaselineWhiteListRequestBodyRuleList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddBaselineWhiteListRequestBodyRuleList struct{}"
	}

	return strings.Join([]string{"AddBaselineWhiteListRequestBodyRuleList", string(data)}, " ")
}
