package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateAgentResponse struct {
	Agent          *Agent `json:"agent,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgentResponse struct{}"
	}

	return strings.Join([]string{"UpdateAgentResponse", string(data)}, " ")
}
