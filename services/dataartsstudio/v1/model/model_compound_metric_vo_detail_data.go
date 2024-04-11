package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompoundMetricVoDetailData 返回数据。
type CompoundMetricVoDetailData struct {
	Value *CompoundMetricVo `json:"value,omitempty"`
}

func (o CompoundMetricVoDetailData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompoundMetricVoDetailData struct{}"
	}

	return strings.Join([]string{"CompoundMetricVoDetailData", string(data)}, " ")
}
