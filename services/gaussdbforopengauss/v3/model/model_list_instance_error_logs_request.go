package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceErrorLogsRequest Request Object
type ListInstanceErrorLogsRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。例如：该参数指定为0，limit指定为10，则只展示第1-10条数据。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为10，不能为负数，最小值为1，最大值为100。例如该参数设定为10，则查询结果最多只显示10条记录。
	Limit *int32 `json:"limit,omitempty"`

	// 开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	StartTime string `json:"start_time"`

	// 结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。只能查询当前时间前30天的错误日志。
	EndTime string `json:"end_time"`
}

func (o ListInstanceErrorLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceErrorLogsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceErrorLogsRequest", string(data)}, " ")
}
