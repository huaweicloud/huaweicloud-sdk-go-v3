package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotifyTrainingJobInformationResponse Response Object
type NotifyTrainingJobInformationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o NotifyTrainingJobInformationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotifyTrainingJobInformationResponse struct{}"
	}

	return strings.Join([]string{"NotifyTrainingJobInformationResponse", string(data)}, " ")
}
