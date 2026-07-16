package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingJobEnginesRequest Request Object
type ShowTrainingJobEnginesRequest struct {
}

func (o ShowTrainingJobEnginesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingJobEnginesRequest struct{}"
	}

	return strings.Join([]string{"ShowTrainingJobEnginesRequest", string(data)}, " ")
}
