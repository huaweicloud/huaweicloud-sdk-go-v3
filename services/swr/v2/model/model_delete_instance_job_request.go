package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceJobRequest Request Object
type DeleteInstanceJobRequest struct {

	// 任务ID
	JobId string `json:"job_id"`
}

func (o DeleteInstanceJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceJobRequest", string(data)}, " ")
}
