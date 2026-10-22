package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HitDetail 检测命中详情
type HitDetail struct {

	// 威胁类型
	RiskType *string `json:"risk_type,omitempty"`

	// 子威胁类型
	SubRiskType *[]string `json:"sub_risk_type,omitempty"`

	// 置信度分数
	Prob *float32 `json:"prob,omitempty"`

	// 命中规则的片段
	Segments *[]SegmentInfo `json:"segments,omitempty"`

	// 额外检测信息，不同风险类型下字段不固定
	ExtraInfo *interface{} `json:"extra_info,omitempty"`
}

func (o HitDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HitDetail struct{}"
	}

	return strings.Join([]string{"HitDetail", string(data)}, " ")
}
