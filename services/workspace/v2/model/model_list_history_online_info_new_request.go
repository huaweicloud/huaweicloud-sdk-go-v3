package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryOnlineInfoNewRequest Request Object
type ListHistoryOnlineInfoNewRequest struct {

	// 查询的起始时间。指定该参数后，返回的结果为此时间之后的所有登录记录。时间格式如：“2016-08-20T21:11Z”。终止时间不为空时，起始时间为必填参数。类型查询优先于时间查询。类型查询和时间查询必须有一个存在。
	StartTime *string `json:"start_time,omitempty"`

	// 查询的结束时间。指定该参数后，返回的结果为此时间之前的所有登录记录。时间格式如：“2016-08-20T21:11Z”。起始时间不为空时，终止时间为必填参数。类型查询优先于时间查询。类型查询和时间查询必须有一个存在。
	EndTime *string `json:"end_time,omitempty"`

	// 查询类型，类型查询优先于时间查询。类型查询和时间查询必须有一个存在。 -MONTH：按月查询。 -WEEK：按周查询。 -DAY：按天查询。
	QueryType *string `json:"query_type,omitempty"`
}

func (o ListHistoryOnlineInfoNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryOnlineInfoNewRequest struct{}"
	}

	return strings.Join([]string{"ListHistoryOnlineInfoNewRequest", string(data)}, " ")
}
