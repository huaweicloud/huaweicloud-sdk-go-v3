package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncLimitDataResponse Response Object
type SyncLimitDataResponse struct {

	// 同步结果，success成功，fail失败
	Result *string `json:"result,omitempty"`

	// 同步数据的实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 同步数据的节点ID，集中式表示主节点ID，分布式表示CN节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 同步的数据记录总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o SyncLimitDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncLimitDataResponse struct{}"
	}

	return strings.Join([]string{"SyncLimitDataResponse", string(data)}, " ")
}
