package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSearchCriteriasRequest Request Object
type CreateSearchCriteriasRequest struct {

	// 租户想查询的日志流所在的日志组的groupid，一般为36位字符串。  缺省值：None  最小长度：36  最大长度：36
	GroupId string `json:"group_id"`

	// 日志流id
	TopicId string `json:"topic_id"`

	Body *CreateSearchCriteriasBody `json:"body,omitempty"`
}

func (o CreateSearchCriteriasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSearchCriteriasRequest struct{}"
	}

	return strings.Join([]string{"CreateSearchCriteriasRequest", string(data)}, " ")
}
