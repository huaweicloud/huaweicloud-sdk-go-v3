package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
