package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCompetitionScoreRequest Request Object
type UpdateCompetitionScoreRequest struct {
	Body *UpdateScoreRequestModel `json:"body,omitempty"`
}

func (o UpdateCompetitionScoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCompetitionScoreRequest struct{}"
	}

	return strings.Join([]string{"UpdateCompetitionScoreRequest", string(data)}, " ")
}
