package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateCompetitionScoreResponse struct {

	// 作品ID
	WorksId        *string `json:"works_id,omitempty" xml:"works_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCompetitionScoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCompetitionScoreResponse struct{}"
	}

	return strings.Join([]string{"CreateCompetitionScoreResponse", string(data)}, " ")
}
