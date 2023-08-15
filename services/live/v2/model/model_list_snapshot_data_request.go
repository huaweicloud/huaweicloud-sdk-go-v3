package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotDataRequest Request Object
type ListSnapshotDataRequest struct {

	// 推流域名。
	PublishDomain *string `json:"publish_domain,omitempty"`

	// 起始时间。日期格式按照ISO8601表示法，并使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。  非整点时间按小时取整，若start_time为2020-08-18T07:20:40Z，则实际查询起始时间为2020-08-18T07:00:00Z。  若start_time为空，则默认查询最近24小时数据。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。日期格式按照ISO8601表示法，并使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。  非整点时间按小时取整，若end_time为2020-08-18T08:20:40Z，则实际查询起始时间为2020-08-18T08:00:00Z。  若参数为空，默认为当前时间。结束时间需大于起始时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListSnapshotDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotDataRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotDataRequest", string(data)}, " ")
}
