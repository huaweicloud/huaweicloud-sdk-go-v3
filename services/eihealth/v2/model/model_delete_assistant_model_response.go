package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssistantModelResponse Response Object
type DeleteAssistantModelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAssistantModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssistantModelResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssistantModelResponse", string(data)}, " ")
}
