package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDatabaseInfo databases 返回参数
type GetDatabaseInfo struct {

	// 逻辑库名称。
	Name string `json:"name"`

	// 逻辑库的工作模式。  - cluster表示逻辑库是拆分模式。 - single表示逻辑库是非拆分模式。
	ShardMode string `json:"shard_mode"`

	// 同一种工作模式下逻辑库分片的数量。
	ShardNumber int32 `json:"shard_number"`

	// 逻辑库状态。
	Status string `json:"status"`

	// 逻辑库的创建时间。
	Created string `json:"created"`

	// 逻辑库关联的RDS实例信息。
	UsedRds []GetDatabaseUsedRds `json:"used_rds"`

	// 单个RDS上的逻辑库分片数。
	ShardUnit int32 `json:"shard_unit"`
}

func (o GetDatabaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDatabaseInfo struct{}"
	}

	return strings.Join([]string{"GetDatabaseInfo", string(data)}, " ")
}
