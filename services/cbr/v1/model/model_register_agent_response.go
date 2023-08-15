package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterAgentResponse Response Object
type RegisterAgentResponse struct {
	Agent          *Agent `json:"agent,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o RegisterAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterAgentResponse struct{}"
	}

	return strings.Join([]string{"RegisterAgentResponse", string(data)}, " ")
}
