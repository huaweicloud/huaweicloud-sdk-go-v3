package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NextMarker 如果存在，则表示还有后续的条目未显示在当前返回体中。请使用该值作为下一次请求的分页标记参数以获得下一页信息。请反复调用该接口直至该字段不存在。
type NextMarker struct {
}

func (o NextMarker) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NextMarker struct{}"
	}

	return strings.Join([]string{"NextMarker", string(data)}, " ")
}
