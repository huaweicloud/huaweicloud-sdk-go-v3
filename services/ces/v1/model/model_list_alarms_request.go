package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAlarmsRequest struct {

	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。
	ContentType string `json:"Content-Type" xml:"Content-Type"`

	// 取值范围(0,100]，默认值为100  用于限制结果数据条数。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 用于标识结果排序方法。  取值说明，默认值为desc。  asc：升序 desc：降序
	Order *string `json:"order,omitempty" xml:"order"`

	// 分页起始值，内容为alarm_id。
	Start *string `json:"start,omitempty" xml:"start"`
}

func (o ListAlarmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmsRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmsRequest", string(data)}, " ")
}
