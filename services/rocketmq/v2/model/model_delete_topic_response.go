package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTopicResponse Response Object
type DeleteTopicResponse struct {

	// 规格变更任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTopicResponse struct{}"
	}

	return strings.Join([]string{"DeleteTopicResponse", string(data)}, " ")
}
