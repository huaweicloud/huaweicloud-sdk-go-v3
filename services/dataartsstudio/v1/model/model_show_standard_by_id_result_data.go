package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStandardByIdResultData 数据标准查询返回data结果。
type ShowStandardByIdResultData struct {
	Value *StandElementValueVoList `json:"value,omitempty"`
}

func (o ShowStandardByIdResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStandardByIdResultData struct{}"
	}

	return strings.Join([]string{"ShowStandardByIdResultData", string(data)}, " ")
}
