package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComplianceItem struct {

	// 节点id
	InstanceId *string `json:"instance_id,omitempty"`

	// 补丁名称
	Title *string `json:"title,omitempty"`

	// 分类
	Classification *string `json:"classification,omitempty"`

	// 严重性级别
	SeverityLevel *string `json:"severity_level,omitempty"`

	// 合规性级别
	ComplianceLevel *string `json:"compliance_level,omitempty"`

	PatchDetail *PatchDetail `json:"patch_detail,omitempty"`
}

func (o ComplianceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComplianceItem struct{}"
	}

	return strings.Join([]string{"ComplianceItem", string(data)}, " ")
}
