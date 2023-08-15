package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRiskConfigsRequest Request Object
type ListRiskConfigsRequest struct {

	// 企业项目ID，查询所有企业项目时填写：all_granted_eps
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 基线名称
	CheckName *string `json:"check_name,omitempty"`

	// 策略组ID
	GroupId *string `json:"group_id,omitempty"`

	// 风险等级，包含如下:   - Security : 安全   - Low : 低危   - Medium : 中危   - High : 高危
	Severity *string `json:"severity,omitempty"`

	// 标准类型，包含如下:   - cn_standard : 等保合规标准   - hw_standard : 华为标准   - qt_standard : 青腾标准
	Standard *string `json:"standard,omitempty"`

	// 服务器id
	HostId *string `json:"host_id,omitempty"`

	// 每页显示数量，默认10
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRiskConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRiskConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListRiskConfigsRequest", string(data)}, " ")
}
