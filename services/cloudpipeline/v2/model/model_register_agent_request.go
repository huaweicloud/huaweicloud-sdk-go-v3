package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RegisterAgentRequest struct {
	Body *SlaveRegister `json:"body,omitempty" xml:"body"`
}

func (o RegisterAgentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterAgentRequest struct{}"
	}

	return strings.Join([]string{"RegisterAgentRequest", string(data)}, " ")
}
