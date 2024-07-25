package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOrchestrationResponse Response Object
type DeleteOrchestrationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteOrchestrationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOrchestrationResponse struct{}"
	}

	return strings.Join([]string{"DeleteOrchestrationResponse", string(data)}, " ")
}
