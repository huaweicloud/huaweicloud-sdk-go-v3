package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type CandidateIntention struct {
	// 候选意图。

	CandidateIntention string `json:"candidate_intention"`
	// 候选意图置信度。

	CandidateConfidence float64 `json:"candidate_confidence"`
}

func (o CandidateIntention) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CandidateIntention struct{}"
	}

	return strings.Join([]string{"CandidateIntention", string(data)}, " ")
}
