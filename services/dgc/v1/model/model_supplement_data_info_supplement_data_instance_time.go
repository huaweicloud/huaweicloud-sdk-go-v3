package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SupplementDataInfoSupplementDataInstanceTime 离散的日期补数据
type SupplementDataInfoSupplementDataInstanceTime struct {

	// 离散的天
	Days *[]string `json:"days,omitempty"`

	// 指定天中的时间段
	TimeOfDay *string `json:"time_of_day,omitempty"`
}

func (o SupplementDataInfoSupplementDataInstanceTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SupplementDataInfoSupplementDataInstanceTime struct{}"
	}

	return strings.Join([]string{"SupplementDataInfoSupplementDataInstanceTime", string(data)}, " ")
}
