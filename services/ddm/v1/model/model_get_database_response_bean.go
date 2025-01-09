package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDatabaseResponseBean database 返回参数
type GetDatabaseResponseBean struct {

	// 逻辑库名称。
	Name string `json:"name"`

	// 逻辑库的创建时间。
	Created int64 `json:"created"`

	// 状态。
	Status string `json:"status"`

	// DDM实例最后更新时间。
	Updated int64 `json:"updated"`

	// 逻辑库分片的详细信息。
	Databases []GetDatabases `json:"databases"`

	// 逻辑库的工作模式。  - cluster表示逻辑库是拆分模式。 - single表示逻辑库是非拆分模式。
	ShardMode string `json:"shard_mode"`

	// 同一种工作模式下逻辑库分片的数量。
	ShardNumber int32 `json:"shard_number"`

	// 单个RDS上的逻辑库分片数。
	ShardUnit int32 `json:"shard_unit"`

	// 连接逻辑库使用的IP:端口。
	DataVips []string `json:"dataVips"`

	// 关联RDS
	UsedRds []GetDatabaseUsedRds `json:"used_rds"`
}

func (o GetDatabaseResponseBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDatabaseResponseBean struct{}"
	}

	return strings.Join([]string{"GetDatabaseResponseBean", string(data)}, " ")
}
