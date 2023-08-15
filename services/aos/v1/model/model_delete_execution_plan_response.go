package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteExecutionPlanResponse Response Object
type DeleteExecutionPlanResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteExecutionPlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteExecutionPlanResponse struct{}"
	}

	return strings.Join([]string{"DeleteExecutionPlanResponse", string(data)}, " ")
}
