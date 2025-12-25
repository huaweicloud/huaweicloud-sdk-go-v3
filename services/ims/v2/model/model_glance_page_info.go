package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GlancePageInfo struct {

	// 分页时，作为查询下一页的marker取值。
	NextMarker string `json:"next_marker"`

	// 当前分页记录数。
	CurrentCount int32 `json:"current_count"`
}

func (o GlancePageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlancePageInfo struct{}"
	}

	return strings.Join([]string{"GlancePageInfo", string(data)}, " ")
}
