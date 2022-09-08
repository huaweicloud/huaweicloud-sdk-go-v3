package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListHistoryWebinarsRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用uuId，若不携带，则后台自动生成
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 语言参数，默认为中文zh-CN, 英文为en-US
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// 查询偏移量,若超过最大数量，则返回最后一页的数据 默认值：0
	Offset *int32 `json:"offset,omitempty"`

	// 查询数量 默认值：0
	Limit *int32 `json:"limit,omitempty"`

	// 搜索条件。支持账号、姓名、手机、邮箱模糊搜索
	SearchKey *string `json:"searchKey,omitempty"`

	// ASC_StartTIME：按会议开始时间升序排序。DSC_StartTIME：按会议开始时间降序排序
	SortType *string `json:"sortType,omitempty"`

	// 开始时间, 格式：yyyy-MM-dd HH:mm。
	StartTime *string `json:"startTime,omitempty"`

	// 结束时间, 格式：yyyy-MM-dd HH:mm。
	EndTime *string `json:"endTime,omitempty"`
}

func (o ListHistoryWebinarsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryWebinarsRequest struct{}"
	}

	return strings.Join([]string{"ListHistoryWebinarsRequest", string(data)}, " ")
}
