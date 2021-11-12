package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCompetitionScoreRequest struct {
	Body *CreateScoresRequestModel `json:"body,omitempty"`
}

func (o CreateCompetitionScoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCompetitionScoreRequest struct{}"
	}

	return strings.Join([]string{"CreateCompetitionScoreRequest", string(data)}, " ")
}
