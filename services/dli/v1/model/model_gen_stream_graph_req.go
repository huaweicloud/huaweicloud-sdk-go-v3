package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenStreamGraphReq 生成SQL静态流图请求参数
type GenStreamGraphReq struct {

	// SQL
	SqlBody string `json:"sql_body"`

	// flink版本。可以为空，也可以填写1.10或1.12。
	FlinkVersion *string `json:"flink_version,omitempty"`

	// CU总数
	CuNumber *int32 `json:"cu_number,omitempty"`

	// 管理单元CU数量
	ManagerCuNumber *int32 `json:"manager_cu_number,omitempty"`

	// 最大并行度
	ParallelNumber *int32 `json:"parallel_number,omitempty"`

	// 单个taskManagerCU数量
	TmCus *int32 `json:"tm_cus,omitempty"`

	// 单个taskManager Slot数量
	TmSlotNum *int32 `json:"tm_slot_num,omitempty"`

	// 算子的配置
	OperatorConfig *string `json:"operator_config,omitempty"`

	// 是否静态资源预估
	StaticEstimator *bool `json:"static_estimator,omitempty"`

	// 作业类型
	JobType *string `json:"job_type,omitempty"`

	// 流图类型
	GraphType *string `json:"graph_type,omitempty"`

	// 每个算子的流量/命中率配置，json格式的字符串。
	StaticEstimatorConfig *string `json:"static_estimator_config,omitempty"`
}

func (o GenStreamGraphReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenStreamGraphReq struct{}"
	}

	return strings.Join([]string{"GenStreamGraphReq", string(data)}, " ")
}
