package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TableColumnStatisticsDescription 表列统计信息的描述信息
type TableColumnStatisticsDescription struct {

	// 最后统计时间
	LastAnalyzedTime *sdktime.SdkTime `json:"last_analyzed_time"`
}

func (o TableColumnStatisticsDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TableColumnStatisticsDescription struct{}"
	}

	return strings.Join([]string{"TableColumnStatisticsDescription", string(data)}, " ")
}
