package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 分页结果信息。
type OpenPageInfo struct {
	// 偏移量。

	Offset int32 `json:"offset"`
	// 查询个数。

	Limit int32 `json:"limit"`
	// 总记录数

	Count int64 `json:"count"`
}

func (o OpenPageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenPageInfo struct{}"
	}

	return strings.Join([]string{"OpenPageInfo", string(data)}, " ")
}
