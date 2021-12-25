package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SqlExecutionReq struct {
	// SQL类型。目前仅支持“presto”类型的SQL。  说明： 只有包含Presto组件的集群才能提交执行presto类型的SQL。

	SqlType string `json:"sql_type"`
	// 待执行的SQL语句。  说明： 目前仅支持执行单条语句，语句中不包含“;”。

	SqlContent string `json:"sql_content"`
	// 执行SQL所在的数据库，默认为default。

	Database *string `json:"database,omitempty"`
	// SQL执行结果的转储文件夹。  说明： 只有select语句才会转储查询的结果。当前仅支持转储到OBS中。

	ArchivePath *string `json:"archive_path,omitempty"`
}

func (o SqlExecutionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlExecutionReq struct{}"
	}

	return strings.Join([]string{"SqlExecutionReq", string(data)}, " ")
}
