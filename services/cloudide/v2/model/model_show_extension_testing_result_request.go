package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExtensionTestingResultRequest Request Object
type ShowExtensionTestingResultRequest struct {

	// 发布商凭证,x-publisher-token和X-Auth-Token必传一个
	XPublisherToken *string `json:"x-publisher-token,omitempty"`

	// 任务id
	TaskId string `json:"task_id"`
}

func (o ShowExtensionTestingResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExtensionTestingResultRequest struct{}"
	}

	return strings.Join([]string{"ShowExtensionTestingResultRequest", string(data)}, " ")
}
