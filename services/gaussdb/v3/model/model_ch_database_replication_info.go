package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChDatabaseReplicationInfo ClickHouse同步信息。
type ChDatabaseReplicationInfo struct {

	// 源数据库。
	SourceDatabase string `json:"source_database"`

	// 目标数据库。
	TargetDatabase string `json:"target_database"`

	// 当前状态。 取值范围： - normal：正常 - abnormal：异常
	Status string `json:"status"`

	// 同步阶段。 取值范围： - wait：等待同步 - failed：同步失败 - incremental：增量同步 - full：全量同步 - other：其他
	Stage string `json:"stage"`

	// 进度百分比。
	Percentage string `json:"percentage"`

	// 追赶阶段。 取值范围： - wait：等待同步 - failed：同步失败 - incremental：增量同步 - full：全量同步 - other：其他
	CatchupStage string `json:"catchup_stage"`
}

func (o ChDatabaseReplicationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChDatabaseReplicationInfo struct{}"
	}

	return strings.Join([]string{"ChDatabaseReplicationInfo", string(data)}, " ")
}
