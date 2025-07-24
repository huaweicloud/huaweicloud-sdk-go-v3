package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunInstancesOptionsPlacement 实例placement，指定服务器策略。不指定placement，按min_count、max_count从空闲服务器中选择。
type RunInstancesOptionsPlacement struct {

	// 指定服务器
	ServerIdSet *[]string `json:"server_id_set,omitempty"`
}

func (o RunInstancesOptionsPlacement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunInstancesOptionsPlacement struct{}"
	}

	return strings.Join([]string{"RunInstancesOptionsPlacement", string(data)}, " ")
}
