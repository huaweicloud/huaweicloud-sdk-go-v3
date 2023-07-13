package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartPausingWorkflowExecutionsResponse Response Object
type StartPausingWorkflowExecutionsResponse struct {

	// 结果返回体
	Result         *interface{} `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o StartPausingWorkflowExecutionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartPausingWorkflowExecutionsResponse struct{}"
	}

	return strings.Join([]string{"StartPausingWorkflowExecutionsResponse", string(data)}, " ")
}
