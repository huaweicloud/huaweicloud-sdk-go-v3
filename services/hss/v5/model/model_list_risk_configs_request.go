package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRiskConfigsRequest struct {

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 基线名称
	CheckType *string `json:"check_type,omitempty" xml:"check_type"`

	// 风险等级，包含如下:   - Security : 安全   - Low : 低危   - Medium : 中危   - High : 高危
	Severity *string `json:"severity,omitempty" xml:"severity"`

	// 标准类型，包含如下:   - cn_standard : 等保合规标准   - hw_standard : 华为标准   - qt_standard : 青腾标准
	Standard *string `json:"standard,omitempty" xml:"standard"`

	// 服务器id
	HostId *string `json:"host_id,omitempty" xml:"host_id"`

	// 每页显示数量，默认10
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`
}

func (o ListRiskConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListRiskConfigsRequest", string(data)}, " ")
}
