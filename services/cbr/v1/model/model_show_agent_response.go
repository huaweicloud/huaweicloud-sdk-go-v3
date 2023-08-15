package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgentResponse Response Object
type ShowAgentResponse struct {
	Agent          *Agent `json:"agent,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgentResponse struct{}"
	}

	return strings.Join([]string{"ShowAgentResponse", string(data)}, " ")
}
