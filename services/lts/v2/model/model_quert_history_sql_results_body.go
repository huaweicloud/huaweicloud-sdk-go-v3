package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuertHistorySqlResultsBody struct {

	// 上次修改时间，时间戳，毫秒数
	LastUseTime *int64 `json:"last_use_time,omitempty"`

	// 历史sql语句
	SqlStatement *string `json:"sql_statement,omitempty"`
}

func (o QuertHistorySqlResultsBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuertHistorySqlResultsBody struct{}"
	}

	return strings.Join([]string{"QuertHistorySqlResultsBody", string(data)}, " ")
}
