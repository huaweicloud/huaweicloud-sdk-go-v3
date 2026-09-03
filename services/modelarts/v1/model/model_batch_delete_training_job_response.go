package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTrainingJobResponse Response Object
type BatchDeleteTrainingJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteTrainingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTrainingJobResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteTrainingJobResponse", string(data)}, " ")
}
