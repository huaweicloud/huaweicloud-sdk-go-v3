package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueCoverRateVo 看板项目的需求覆盖率
type IssueCoverRateVo struct {
	Epic *CoverRateVo `json:"epic,omitempty"`

	Feature *CoverRateVo `json:"feature,omitempty"`

	Story *CoverRateVo `json:"story,omitempty"`

	Summary *CoverRateVo `json:"summary,omitempty"`
}

func (o IssueCoverRateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueCoverRateVo struct{}"
	}

	return strings.Join([]string{"IssueCoverRateVo", string(data)}, " ")
}
