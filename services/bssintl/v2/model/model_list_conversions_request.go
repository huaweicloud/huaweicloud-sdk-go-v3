package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConversionsRequest Request Object
type ListConversionsRequest struct {

	// 语言，字段预留。默认zh_CN，枚举：zh_CN：中文 en_US：英文
	XLanguage *string `json:"X-Language,omitempty"`

	// 度量类型。1：货币 2：时长 3：流量 4：数量 7：容量 9：行数 10：周期 11：频率 12：个数 16：带宽速率 19：带宽速率（1000进制） 20：性能测试用量 27：核数*时长 28：内存*时长 29：IOPS*时长 30：吞吐量*时长 31：个/时长 40：流量（1000进制） 41：1K Tokens 108：数量*时长。此参数不携带或携带值为空或携带值为null时，不作为筛选条件；不支持携带值为空串。
	MeasureType *int32 `json:"measure_type,omitempty"`
}

func (o ListConversionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConversionsRequest struct{}"
	}

	return strings.Join([]string{"ListConversionsRequest", string(data)}, " ")
}
