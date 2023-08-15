package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnregisterAgentResponse Response Object
type UnregisterAgentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UnregisterAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnregisterAgentResponse struct{}"
	}

	return strings.Join([]string{"UnregisterAgentResponse", string(data)}, " ")
}
