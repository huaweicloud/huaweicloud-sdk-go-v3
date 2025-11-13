package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAssistantModelResponse Response Object
type UpdateAssistantModelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAssistantModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAssistantModelResponse struct{}"
	}

	return strings.Join([]string{"UpdateAssistantModelResponse", string(data)}, " ")
}
