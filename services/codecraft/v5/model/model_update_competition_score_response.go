package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCompetitionScoreResponse Response Object
type UpdateCompetitionScoreResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateCompetitionScoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCompetitionScoreResponse struct{}"
	}

	return strings.Join([]string{"UpdateCompetitionScoreResponse", string(data)}, " ")
}
