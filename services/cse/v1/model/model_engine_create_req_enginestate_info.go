package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EngineCreateReqEnginestateInfo 引擎状态信息
type EngineCreateReqEnginestateInfo struct {

	// 集群
	Cluster *bool `json:"cluster,omitempty"`

	// 双子集群
	TwinClusters *bool `json:"twinClusters,omitempty"`

	// 单引擎
	SingleEngine *bool `json:"singleEngine,omitempty"`
}

func (o EngineCreateReqEnginestateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineCreateReqEnginestateInfo struct{}"
	}

	return strings.Join([]string{"EngineCreateReqEnginestateInfo", string(data)}, " ")
}
