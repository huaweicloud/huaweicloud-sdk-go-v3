package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupplementDataRespSupplementDataInstanceTime 支持补离散时间任务
type SupplementDataRespSupplementDataInstanceTime struct {

	// 支持离散的天
	Days *[]string `json:"days,omitempty"`

	// 一天中的时间段
	TimeOfDay *string `json:"time_of_day,omitempty"`
}

func (o SupplementDataRespSupplementDataInstanceTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupplementDataRespSupplementDataInstanceTime struct{}"
	}

	return strings.Join([]string{"SupplementDataRespSupplementDataInstanceTime", string(data)}, " ")
}
