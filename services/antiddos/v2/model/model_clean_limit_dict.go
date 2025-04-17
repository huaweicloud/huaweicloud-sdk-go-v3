package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CleanLimitDict 连接数限制列表
type CleanLimitDict struct {

	// 清洗时访问限制分段ID，取值范围：1：10M;2：30M;3：50M;4：70M;5：100M;6：150M;7：200M;8：250M;9：300M;10：500M;11：800M;88：1000M;99：默认防护。
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
