package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EngineDistributionInfo EngineDistributionInfo对象
type EngineDistributionInfo struct {

	// 数据库类型
	EngineType *string `json:"engine_type,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 实例信息
	InstanceInfos *[]DistributionInstanceInfo `json:"instance_infos,omitempty"`
}

func (o EngineDistributionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineDistributionInfo struct{}"
	}

	return strings.Join([]string{"EngineDistributionInfo", string(data)}, " ")
}
