package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesignCompoundMetricResultData 创建复合指标的返回结果，成功的个数。
type CreateDesignCompoundMetricResultData struct {
	Value *CompoundMetricVo `json:"value,omitempty"`
}

func (o CreateDesignCompoundMetricResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesignCompoundMetricResultData struct{}"
	}

	return strings.Join([]string{"CreateDesignCompoundMetricResultData", string(data)}, " ")
}
