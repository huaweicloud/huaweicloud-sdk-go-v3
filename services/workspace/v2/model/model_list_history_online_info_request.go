package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryOnlineInfoRequest Request Object
type ListHistoryOnlineInfoRequest struct {

	// 查询的起始时间。指定该参数后，返回的结果为此时间之后的所有登录记录。时间格式如：“2016-08-20T21:11Z”。终止时间不为空时，起始时间为必填参数
	StartTime *string `json:"start_time,omitempty"`

	// 查询的终止时间。指定该参数后，返回的结果为此时间之前的所有登录记录。时间格式如：“2016-08-20T21:11Z”。起始时间不为空时，终止时间为必填参数
	EndTime *string `json:"end_time,omitempty"`

	// 查询类型，合法取值有三个:MONTH按月查询 WEEK：按周查询DAY：按天查询
	QueryType *string `json:"query_type,omitempty"`

	// 客户端所在操作系统时间的小时数
	ClientHour *int32 `json:"client_hour,omitempty"`
}

func (o ListHistoryOnlineInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryOnlineInfoRequest struct{}"
	}

	return strings.Join([]string{"ListHistoryOnlineInfoRequest", string(data)}, " ")
}
