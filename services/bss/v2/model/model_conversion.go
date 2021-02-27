package model

import (
	"encoding/json"

	"strings"
)

type Conversion struct {
	// 度量单位ID。 例如：10表示GB。
	MeasureId *int32 `json:"measure_id,omitempty"`
	// 转换后的度量单位ID。 例如：11表示MB。
	RefMeasureId *int32 `json:"ref_measure_id,omitempty"`
	// 度量单位和转换后的度量单位之间的转换比率。 例如： 度量单位为GB，转换度量单位为MB时，转换比率为1024，两者之间的转换公式为：1GB=1024MB。
	ConversionRatio *int64 `json:"conversion_ratio,omitempty"`
	// 度量类型。 1：货币2：时长3：流量4：数量7：容量9：行数10：周期11：频率12：个数16：带宽速率17：容量时长18：查询速率19：带宽速率（1000进制）20：性能测试用量21：面积22：视频23：吞吐量25：测试类型
	MeasureType *int32 `json:"measure_type,omitempty"`
}

func (o Conversion) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Conversion struct{}"
	}

	return strings.Join([]string{"Conversion", string(data)}, " ")
}
