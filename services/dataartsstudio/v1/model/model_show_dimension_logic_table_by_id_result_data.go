package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDimensionLogicTableByIdResultData data，统一的返回结果的最外层数据结构。
type ShowDimensionLogicTableByIdResultData struct {
	Value *DimensionLogicTableVo `json:"value,omitempty"`
}

func (o ShowDimensionLogicTableByIdResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDimensionLogicTableByIdResultData struct{}"
	}

	return strings.Join([]string{"ShowDimensionLogicTableByIdResultData", string(data)}, " ")
}
