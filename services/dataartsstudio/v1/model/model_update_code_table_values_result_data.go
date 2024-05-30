package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCodeTableValuesResultData data，统一的返回结果的最外层数据结构。
type UpdateCodeTableValuesResultData struct {

	// 码表字段列表信息。
	Value *[]CodeTableFieldVo `json:"value,omitempty"`
}

func (o UpdateCodeTableValuesResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCodeTableValuesResultData struct{}"
	}

	return strings.Join([]string{"UpdateCodeTableValuesResultData", string(data)}, " ")
}
