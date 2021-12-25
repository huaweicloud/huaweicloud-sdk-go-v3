package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoDailyCodeline struct {
	// 每日增加代码行

	Additions *int32 `json:"additions,omitempty"`
	// 日期

	Date *string `json:"date,omitempty"`
	// 每日删除代码行

	Deletions *int32 `json:"deletions,omitempty"`
}

func (o RepoDailyCodeline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoDailyCodeline struct{}"
	}

	return strings.Join([]string{"RepoDailyCodeline", string(data)}, " ")
}
