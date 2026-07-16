package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrainingExperimentResponse Response Object
type DeleteTrainingExperimentResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTrainingExperimentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrainingExperimentResponse struct{}"
	}

	return strings.Join([]string{"DeleteTrainingExperimentResponse", string(data)}, " ")
}
