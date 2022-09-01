package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRiskConfigDetailRequest struct {

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 基线名称
	CheckType string `json:"check_type" xml:"check_type"`

	// 标准类型，包含如下: - cn_standard : 等保合规标准 - hw_standard : 华为标准 - qt_standard : 青腾标准
	Standard string `json:"standard" xml:"standard"`

	// 主机ID，不赋值时，查租户所有主机
	HostId *string `json:"host_id,omitempty" xml:"host_id"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`
}

func (o ShowRiskConfigDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRiskConfigDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowRiskConfigDetailRequest", string(data)}, " ")
}
