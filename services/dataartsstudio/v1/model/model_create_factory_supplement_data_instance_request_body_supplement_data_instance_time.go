package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFactorySupplementDataInstanceRequestBodySupplementDataInstanceTime 离散的日期补数据
type CreateFactorySupplementDataInstanceRequestBodySupplementDataInstanceTime struct {

	// 离散的天
	Days *[]string `json:"days,omitempty"`

	// 指定天中的时间段
	TimeOfDay *string `json:"time_of_day,omitempty"`
}

func (o CreateFactorySupplementDataInstanceRequestBodySupplementDataInstanceTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFactorySupplementDataInstanceRequestBodySupplementDataInstanceTime struct{}"
	}

	return strings.Join([]string{"CreateFactorySupplementDataInstanceRequestBodySupplementDataInstanceTime", string(data)}, " ")
}
