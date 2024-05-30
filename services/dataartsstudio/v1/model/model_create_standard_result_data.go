package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateStandardResultData 数据标准创建返回data结果。
type CreateStandardResultData struct {
	Value *StandElementValueVoList `json:"value,omitempty"`
}

func (o CreateStandardResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateStandardResultData struct{}"
	}

	return strings.Join([]string{"CreateStandardResultData", string(data)}, " ")
}
