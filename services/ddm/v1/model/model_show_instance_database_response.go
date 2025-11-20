package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceDatabaseResponse Response Object
type ShowInstanceDatabaseResponse struct {

	// 逻辑库分片信息。
	Shards *[]Shards `json:"shards,omitempty"`

	// 逻辑库状态。
	Status *string `json:"status,omitempty"`

	// 创建时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Created *string `json:"created,omitempty"`

	// 更新时间，格式为\"yyyy-MM-ddTHH:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	Updated *string `json:"updated,omitempty"`

	// 逻辑库名称。
	Name *string `json:"name,omitempty"`

	// 拆分模式。
	ShardMode *string `json:"shard_mode,omitempty"`

	// 逻辑库分片数。
	ShardNumber *int32 `json:"shard_number,omitempty"`

	// 关联的后端DN信息。
	DataNodes      *[]DataNodes `json:"data_nodes,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowInstanceDatabaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceDatabaseResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceDatabaseResponse", string(data)}, " ")
}
