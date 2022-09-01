package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StopJobResponse struct {

	// 作业运行信息，请参见submission参数说明
	Submissions    *[]StartJobSubmission `json:"submissions,omitempty" xml:"submissions"`
	HttpStatusCode int                   `json:"-"`
}

func (o StopJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobResponse struct{}"
	}

	return strings.Join([]string{"StopJobResponse", string(data)}, " ")
}
