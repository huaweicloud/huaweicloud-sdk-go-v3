package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConversionsRequest Request Object
type ListConversionsRequest struct {

	// 语言。zh_CN：中文en_US：英文缺省为zh_CN。
	XLanguage *string `json:"X-Language,omitempty"`

	// 度量类型。1：货币2：时长3：流量4：数量6：长度7：容量9：行数10：周期11：频率12：个数15：容量*时长16：带宽速率17：容量时长18：查询速率19：带宽速率（1000进制）20：性能测试用量21：面积22：视频23：吞吐量24：分辨率26：通用资源包抵扣单位27：核数*时长28：内存*时长29：IOPS*时长30：吞吐量*时长31：个/时长40：流量（1000进制）41：1K Tokens82：帧98：缓存带宽x时长此参数不携带或携带值为空或携带值为null时，不作为筛选条件。
	MeasureType *int32 `json:"measure_type,omitempty"`
}

func (o ListConversionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConversionsRequest struct{}"
	}

	return strings.Join([]string{"ListConversionsRequest", string(data)}, " ")
}
