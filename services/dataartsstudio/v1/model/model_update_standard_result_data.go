package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStandardResultData 数据标准更新返回data结果。
type UpdateStandardResultData struct {
	Value *StandElementValueVoList `json:"value,omitempty"`
}

func (o UpdateStandardResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStandardResultData struct{}"
	}

	return strings.Join([]string{"UpdateStandardResultData", string(data)}, " ")
}
