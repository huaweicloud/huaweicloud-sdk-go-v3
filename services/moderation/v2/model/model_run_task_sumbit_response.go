package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunTaskSumbitResponse Response Object
type RunTaskSumbitResponse struct {
	Result         *TaskSumbitResponseResult `json:"result,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o RunTaskSumbitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTaskSumbitResponse struct{}"
	}

	return strings.Join([]string{"RunTaskSumbitResponse", string(data)}, " ")
}
