package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartInstanceResizeCheckJobResponse Response Object
type StartInstanceResizeCheckJobResponse struct {

	// 任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartInstanceResizeCheckJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceResizeCheckJobResponse struct{}"
	}

	return strings.Join([]string{"StartInstanceResizeCheckJobResponse", string(data)}, " ")
}
