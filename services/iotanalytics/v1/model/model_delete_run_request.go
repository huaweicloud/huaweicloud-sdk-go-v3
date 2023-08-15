package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRunRequest Request Object
type DeleteRunRequest struct {

	// 作业ID。
	JobId string `json:"job_id"`

	// 作业运行ID。
	RunId string `json:"run_id"`
}

func (o DeleteRunRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRunRequest struct{}"
	}

	return strings.Join([]string{"DeleteRunRequest", string(data)}, " ")
}
