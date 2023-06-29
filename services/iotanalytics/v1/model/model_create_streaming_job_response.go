package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStreamingJobResponse Response Object
type CreateStreamingJobResponse struct {

	// 作业ID
	JobId *string `json:"job_id,omitempty"`

	// 作业错误详情
	CheckInfo      map[string]interface{} `json:"check_info,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o CreateStreamingJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStreamingJobResponse struct{}"
	}

	return strings.Join([]string{"CreateStreamingJobResponse", string(data)}, " ")
}
