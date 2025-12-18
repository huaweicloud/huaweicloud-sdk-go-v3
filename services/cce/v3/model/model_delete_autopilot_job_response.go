package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAutopilotJobResponse Response Object
type DeleteAutopilotJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAutopilotJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAutopilotJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteAutopilotJobResponse", string(data)}, " ")
}
