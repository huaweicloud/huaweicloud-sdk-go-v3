package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactLogicTableByIdResultData data，统一的返回结果的最外层数据结构。
type ShowFactLogicTableByIdResultData struct {
	Value *FactLogicTableVo `json:"value,omitempty"`
}

func (o ShowFactLogicTableByIdResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactLogicTableByIdResultData struct{}"
	}

	return strings.Join([]string{"ShowFactLogicTableByIdResultData", string(data)}, " ")
}
