package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 起止时间范围
type TimeSpan struct {
	// 起始时间

	From string `json:"from"`
	// 结束时间

	To string `json:"to"`
}

func (o TimeSpan) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeSpan struct{}"
	}

	return strings.Join([]string{"TimeSpan", string(data)}, " ")
}
