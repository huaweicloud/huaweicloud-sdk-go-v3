package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountAllModelsResultData data，统一的返回结果的最外层数据结构。
type CountAllModelsResultData struct {
	Value *AllModelStatisticVo `json:"value,omitempty"`
}

func (o CountAllModelsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountAllModelsResultData struct{}"
	}

	return strings.Join([]string{"CountAllModelsResultData", string(data)}, " ")
}
