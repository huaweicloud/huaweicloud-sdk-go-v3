package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesignCompoundMetricResultData 更新复合指标的返回结果，成功的个数。
type UpdateDesignCompoundMetricResultData struct {
	Value *CompoundMetricVo `json:"value,omitempty"`
}

func (o UpdateDesignCompoundMetricResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesignCompoundMetricResultData struct{}"
	}

	return strings.Join([]string{"UpdateDesignCompoundMetricResultData", string(data)}, " ")
}
