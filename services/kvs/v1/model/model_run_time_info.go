package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RunTimeInfo struct {
	TableInfo *TableInfo `bson:"table_info"`

	// 索引状态。
	LocalSecondaryIndexInfos *[]SecondaryIndexInfo `bson:"local_secondary_index_infos,omitempty"`

	// 全局二级索引运行态。
	GlobalSecondaryIndexInfos *[]GlobalSecondaryIndexInfo `bson:"global_secondary_index_infos,omitempty"`
}

func (o RunTimeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTimeInfo struct{}"
	}

	return strings.Join([]string{"RunTimeInfo", string(data)}, " ")
}
