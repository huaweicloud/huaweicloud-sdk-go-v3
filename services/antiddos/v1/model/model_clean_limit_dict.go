package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 连接数限制列表
type CleanLimitDict struct {
	// 清洗时访问限制分段ID

	CleaningAccessPosId int64 `json:"cleaning_access_pos_id"`
	// 单一源IP新建连接个数

	NewConnectionLimited int64 `json:"new_connection_limited"`
	// 单一源IP连接数总个数

	TotalConnectionLimited int64 `json:"total_connection_limited"`
}

func (o CleanLimitDict) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CleanLimitDict struct{}"
	}

	return strings.Join([]string{"CleanLimitDict", string(data)}, " ")
}
