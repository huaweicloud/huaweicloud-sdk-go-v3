package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopJobResponse struct {
	// 作业运行信息，请参见submission参数说明

	Submissions    *[]StartJobSubmission `json:"submissions,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o StopJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobResponse struct{}"
	}

	return strings.Join([]string{"StopJobResponse", string(data)}, " ")
}
