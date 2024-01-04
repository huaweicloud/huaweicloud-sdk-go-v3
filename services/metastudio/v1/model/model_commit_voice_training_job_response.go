package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommitVoiceTrainingJobResponse Response Object
type CommitVoiceTrainingJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CommitVoiceTrainingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitVoiceTrainingJobResponse struct{}"
	}

	return strings.Join([]string{"CommitVoiceTrainingJobResponse", string(data)}, " ")
}
