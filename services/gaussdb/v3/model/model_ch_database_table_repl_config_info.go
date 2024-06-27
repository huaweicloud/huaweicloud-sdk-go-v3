package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChDatabaseTableReplConfigInfo 表配置信息。
type ChDatabaseTableReplConfigInfo struct {

	// 表同步类型。 取值范围： - white_list：白名单，此时表范围不能为空。 - black_list：黑名单，此时表范围为空则选择所有表。
	ReplType string `json:"repl_type"`

	// 白名单或黑名单的表范围。
	Tables []string `json:"tables"`
}

func (o ChDatabaseTableReplConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChDatabaseTableReplConfigInfo struct{}"
	}

	return strings.Join([]string{"ChDatabaseTableReplConfigInfo", string(data)}, " ")
}
