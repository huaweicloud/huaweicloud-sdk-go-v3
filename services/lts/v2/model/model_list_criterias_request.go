package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCriteriasRequest Request Object
type ListCriteriasRequest struct {

	// 租户想查询的日志流所在的日志组的groupid，一般为36位字符串。  缺省值：None  最小长度：36  最大长度：36
	GroupId string `json:"group_id"`

	// 日志流id
	TopicId string `json:"topic_id"`

	// 原始日志：ORIGINALLOG 可视化日志: VISUALIZATION
	SearchType *string `json:"search_type,omitempty"`
}

func (o ListCriteriasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCriteriasRequest struct{}"
	}

	return strings.Join([]string{"ListCriteriasRequest", string(data)}, " ")
}
