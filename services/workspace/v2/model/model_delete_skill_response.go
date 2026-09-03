package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSkillResponse Response Object
type DeleteSkillResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteSkillResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSkillResponse struct{}"
	}

	return strings.Join([]string{"DeleteSkillResponse", string(data)}, " ")
}
