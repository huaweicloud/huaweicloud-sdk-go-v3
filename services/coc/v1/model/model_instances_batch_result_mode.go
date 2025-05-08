package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstancesBatchResultMode 获取分批结果返回体
type InstancesBatchResultMode struct {

	// 批次Id
	BatchIndex int32 `json:"batch_index"`

	// 当前批次内机器
	TargetInstances []ResourceInstance `json:"target_instances"`
}

func (o InstancesBatchResultMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesBatchResultMode struct{}"
	}

	return strings.Join([]string{"InstancesBatchResultMode", string(data)}, " ")
}
