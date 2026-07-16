package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrainingJobResponse Response Object
type DeleteTrainingJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTrainingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrainingJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteTrainingJobResponse", string(data)}, " ")
}
