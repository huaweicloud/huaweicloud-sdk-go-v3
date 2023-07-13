package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTopicOrBatchDeleteTopicResponse Response Object
type CreateTopicOrBatchDeleteTopicResponse struct {

	// 主题名称。
	Id *string `json:"id,omitempty"`

	// 删除主题任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateTopicOrBatchDeleteTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTopicOrBatchDeleteTopicResponse struct{}"
	}

	return strings.Join([]string{"CreateTopicOrBatchDeleteTopicResponse", string(data)}, " ")
}
