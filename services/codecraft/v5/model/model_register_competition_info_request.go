package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterCompetitionInfoRequest Request Object
type RegisterCompetitionInfoRequest struct {
	Body *RegisterInfoRequestModel `json:"body,omitempty"`
}

func (o RegisterCompetitionInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterCompetitionInfoRequest struct{}"
	}

	return strings.Join([]string{"RegisterCompetitionInfoRequest", string(data)}, " ")
}
