package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelationInfo struct {

	// 需求id
	DrNumber string `json:"dr_number"`

	// 用例uri
	TestCaseUri *string `json:"test_case_uri,omitempty"`

	// 资源类型
	RelateType string `json:"relate_type"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 来源系统
	SourceSystem *string `json:"source_system,omitempty"`

	// 关联资源编号
	AssociationNumber *string `json:"association_number,omitempty"`

	// 逻辑region，外部使用公有云实际区域，内部使用默认值
	Region *string `json:"region,omitempty"`
}

func (o RelationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelationInfo struct{}"
	}

	return strings.Join([]string{"RelationInfo", string(data)}, " ")
}
