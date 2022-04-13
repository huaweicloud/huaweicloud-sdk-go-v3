package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartJobResponse struct {
	// 作业运行信息，请参见submission参数说明

	Submissions    *[]StartJobSubmission `json:"submissions,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o StartJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartJobResponse struct{}"
	}

	return strings.Join([]string{"StartJobResponse", string(data)}, " ")
}
