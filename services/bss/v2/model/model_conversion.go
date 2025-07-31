package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Conversion struct {

	// 度量单位ID。 例如：10表示GB。
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 转换后的度量单位ID。 例如：11表示MB。
	RefMeasureId *int32 `json:"ref_measure_id,omitempty"`

	// 度量单位和转换后的度量单位之间的转换比率。 例如： 度量单位为GB，转换度量单位为MB时，转换比率为1024，两者之间的转换公式为：1GB=1024MB。
	ConversionRatio *int64 `json:"conversion_ratio,omitempty"`

	// 度量类型。1：货币 2：时长 3：流量 4：数量 7：容量 9：行数 10：周期 11：频率 12：个数 16：带宽速率 19：带宽速率（1000进制）20：性能测试用量 27：核数*时长 28：内存*时长 29：IOPS*时长 30：吞吐量*时长 31：个/时长 40：流量（1000进制）41：1K Tokens 98：缓存带宽x时长 104：tokens 108：数量*时长
	MeasureType *int32 `json:"measure_type,omitempty"`
}

func (o Conversion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Conversion struct{}"
	}

	return strings.Join([]string{"Conversion", string(data)}, " ")
}
