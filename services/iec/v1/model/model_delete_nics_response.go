package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteNicsResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteNicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNicsResponse struct{}"
	}

	return strings.Join([]string{"DeleteNicsResponse", string(data)}, " ")
}
