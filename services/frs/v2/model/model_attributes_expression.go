package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 人脸表情，包括中性、高兴、害怕、惊讶、伤心、生气、厌恶。
type AttributesExpression struct {
	// 人脸表情类型： • neutral：中性 • happy：高兴 • fear：害怕 • surprise：惊讶 • sad：伤心 • angry：生气 • disgust：厌恶 • unknown：图片质量问题导致未识别

	Type *string `json:"type,omitempty"`
	// 表情置信度，取值范围[0-1]。

	Probability *float64 `json:"probability,omitempty"`
}

func (o AttributesExpression) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttributesExpression struct{}"
	}

	return strings.Join([]string{"AttributesExpression", string(data)}, " ")
}
