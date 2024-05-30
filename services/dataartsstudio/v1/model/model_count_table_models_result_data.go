package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountTableModelsResultData data，统一的返回结果的最外层数据结构。
type CountTableModelsResultData struct {
	Value *TableModelStatisticVo `json:"value,omitempty"`
}

func (o CountTableModelsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountTableModelsResultData struct{}"
	}

	return strings.Join([]string{"CountTableModelsResultData", string(data)}, " ")
}
