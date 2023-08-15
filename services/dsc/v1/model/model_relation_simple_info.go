package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelationSimpleInfo struct {

	// 关系信息ID
	Id *string `json:"id,omitempty"`

	// 关系信息名称
	Name *string `json:"name,omitempty"`

	// 关系信息路径
	Path *string `json:"path,omitempty"`

	// 风险等级
	RiskLevel *int32 `json:"risk_level,omitempty"`

	// 关系信息类型
	Type *string `json:"type,omitempty"`
}

func (o RelationSimpleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelationSimpleInfo struct{}"
	}

	return strings.Join([]string{"RelationSimpleInfo", string(data)}, " ")
}
