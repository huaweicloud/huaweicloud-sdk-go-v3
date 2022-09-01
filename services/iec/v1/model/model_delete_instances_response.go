package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteInstancesResponse struct {

	// 边缘任务ID。
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstancesResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstancesResponse", string(data)}, " ")
}
