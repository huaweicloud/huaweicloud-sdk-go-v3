package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVoiceTrainingJobResponse Response Object
type DeleteVoiceTrainingJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVoiceTrainingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVoiceTrainingJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteVoiceTrainingJobResponse", string(data)}, " ")
}
