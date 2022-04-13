package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询结果元数据统计个数。
type TotalMetaData struct {
	// 总条数。

	Total *int32 `json:"total,omitempty"`
}

func (o TotalMetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TotalMetaData struct{}"
	}

	return strings.Join([]string{"TotalMetaData", string(data)}, " ")
}
