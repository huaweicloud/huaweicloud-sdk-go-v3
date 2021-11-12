package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteInstanceResponse struct {
	// 实例删除的任务id。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstanceResponse", string(data)}, " ")
}
