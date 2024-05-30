package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDimensionByIdResultData data，统一的返回结果的最外层数据结构。
type ShowDimensionByIdResultData struct {
	Value *DimensionVo `json:"value,omitempty"`
}

func (o ShowDimensionByIdResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDimensionByIdResultData struct{}"
	}

	return strings.Join([]string{"ShowDimensionByIdResultData", string(data)}, " ")
}
