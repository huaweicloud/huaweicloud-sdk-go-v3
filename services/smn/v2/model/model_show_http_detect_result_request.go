package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpDetectResultRequest Request Object
type ShowHttpDetectResultRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn string `json:"topic_urn"`

	// http探测任务id
	TaskId string `json:"task_id"`
}

func (o ShowHttpDetectResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpDetectResultRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpDetectResultRequest", string(data)}, " ")
}
