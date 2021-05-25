package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListConversionsRequest struct {
	// |忽略大小写，默认 zh_CN：中文 en_US：英文|

	XLanguage *string `json:"X-Language,omitempty"`
	// 度量类型。1：货币2：时长3：流量4：数量7：容量9：行数10：周期11：频率12：个数16：带宽速率17：容量时长18：查询速率19：带宽速率（1000进制）20：性能测试用量21：面积22：视频23：吞吐量25：测试类型

	MeasureType *int32 `json:"measure_type,omitempty"`
}

func (o ListConversionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListConversionsRequest struct{}"
	}

	return strings.Join([]string{"ListConversionsRequest", string(data)}, " ")
}
