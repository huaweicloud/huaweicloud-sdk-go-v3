package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmsRequest Request Object
type ListAlarmsRequest struct {

	// 取值范围(0,100]，默认值为100  用于限制结果数据条数。
	Limit *int32 `json:"limit,omitempty"`

	// 用于标识结果排序方法。  取值说明，默认值为desc。  asc：升序 desc：降序
	Order *string `json:"order,omitempty"`

	// 分页起始值，内容为alarm_id。
	Start *string `json:"start,omitempty"`
}

func (o ListAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmsRequest", string(data)}, " ")
}
