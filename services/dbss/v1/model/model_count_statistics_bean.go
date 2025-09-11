package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CountStatisticsBean struct {

	// 周期
	Period *string `json:"period,omitempty"`

	// SQL数量
	SqlNum *int64 `json:"sql_num,omitempty"`
}

func (o CountStatisticsBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountStatisticsBean struct{}"
	}

	return strings.Join([]string{"CountStatisticsBean", string(data)}, " ")
}
