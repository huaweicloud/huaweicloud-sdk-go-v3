package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InjectionStatisticsBean struct {

	// 注入数量
	InjectionNum *int64 `json:"injection_num,omitempty"`

	// 周期
	Period *string `json:"period,omitempty"`
}

func (o InjectionStatisticsBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InjectionStatisticsBean struct{}"
	}

	return strings.Join([]string{"InjectionStatisticsBean", string(data)}, " ")
}
