package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRestoreTimeRequest Request Object
type ListRestoreTimeRequest struct {

	// 实例Id，可以调用[5.3.3 查询实例列表和详情](x-wc://file=zh-cn_topic_0000001397299481.xml)接口获取。如果未申请实例，可以调用[5.3.1 创建实例](x-wc://file=zh-cn_topic_0000001397139461.xml)接口创建。
	InstanceId string `json:"instance_id"`

	// 查询的可恢复时间段的开始时间点，为yyyy-mm-ddThh:mm:ssZ字符串格式，T指某个时间的开始，Z指时区偏移量。  [例如北京时间偏移显示为+0800。默认值为当前查询时间的前一天。]
	StartTime *string `json:"start_time,omitempty"`

	// 查询的可恢复时间段的结束时间点，为yyyy-mm-ddThh:mm:ssZ字符串格式，T指某个时间的开始，Z指时区偏移量。  [例如北京时间偏移显示为+0800。默认值为当前查询时间。]
	EndTime *string `json:"end_time,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录，默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量上限值，取值范围为0~1000，默认值为1000。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRestoreTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreTimeRequest struct{}"
	}

	return strings.Join([]string{"ListRestoreTimeRequest", string(data)}, " ")
}
