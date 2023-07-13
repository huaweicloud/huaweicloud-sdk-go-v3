package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSqlStatementsResponse Response Object
type ExportSqlStatementsResponse struct {

	// 全量SQL集合。当集合为空时，说明SQL已全部导出。
	Statements *[]FullSql `json:"statements,omitempty"`

	// 获取下一页所需的标识符。marker仅在3分钟内有效。
	NextMarker     *string `json:"next_marker,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportSqlStatementsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSqlStatementsResponse struct{}"
	}

	return strings.Join([]string{"ExportSqlStatementsResponse", string(data)}, " ")
}
