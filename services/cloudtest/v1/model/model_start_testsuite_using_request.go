package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartTestsuiteUsingRequest Request Object
type StartTestsuiteUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	// 任务id
	SuiteId string `json:"suite_id"`

	Body *TaskActionParamsV5 `json:"body,omitempty"`
}

func (o StartTestsuiteUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartTestsuiteUsingRequest struct{}"
	}

	return strings.Join([]string{"StartTestsuiteUsingRequest", string(data)}, " ")
}
