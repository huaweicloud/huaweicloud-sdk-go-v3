package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RegisterCompetitionInfoResponse struct {
	// 是否允许提交作品，true-允许，false-不允许

	IsPermitted *bool `json:"is_permitted,omitempty"`
	// 团队ID

	TeamId         *string `json:"team_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterCompetitionInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterCompetitionInfoResponse struct{}"
	}

	return strings.Join([]string{"RegisterCompetitionInfoResponse", string(data)}, " ")
}
