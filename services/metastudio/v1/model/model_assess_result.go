package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssessResult 音频质量评估结果
type AssessResult struct {
	Speed *AssessProperty `json:"speed,omitempty"`

	Sound *AssessProperty `json:"sound,omitempty"`

	Snr *AssessProperty `json:"snr,omitempty"`

	Reverb *AssessProperty `json:"reverb,omitempty"`

	DnsmosOvrl *AssessProperty `json:"dnsmos_ovrl,omitempty"`

	DnsmosBak *AssessProperty `json:"dnsmos_bak,omitempty"`

	// 综合建议
	Suggestion *string `json:"suggestion,omitempty"`
}

func (o AssessResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssessResult struct{}"
	}

	return strings.Join([]string{"AssessResult", string(data)}, " ")
}
