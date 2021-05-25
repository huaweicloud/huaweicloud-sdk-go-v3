package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAlarmsRequest struct {
	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。

	ContentType string `json:"Content-Type"`
	// 取值范围(0,100]，默认值为100  用于限制结果数据条数。

	Limit *int32 `json:"limit,omitempty"`
	// 用于标识结果排序方法。  取值说明，默认值为desc。  asc：升序 desc：降序

	Order *string `json:"order,omitempty"`
	// 分页起始值，内容为alarm_id。

	Start *string `json:"start,omitempty"`
}

func (o ListAlarmsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAlarmsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmsRequest", string(data)}, " ")
}
