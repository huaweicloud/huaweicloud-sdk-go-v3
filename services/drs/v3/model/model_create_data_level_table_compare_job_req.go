package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataLevelTableCompareJobReq 创建数据级对比任务请求
type CreateDataLevelTableCompareJobReq struct {

	// 对比类型。 - lines：行数对比 - contents：内容对比 - random：抽样对比，当前仅支持gaussdbv5、gaussdbv5-to-postgresql、gaussdbv5ha-to-postgresql链路。
	Type string `json:"type"`

	// 对比任务启动时间，时间戳格式，取值为空代表立即启动。
	StartTime *string `json:"start_time,omitempty"`

	// 数据级对比模式，取值为空时需要在compare_object或者compare_object_with_token传对象信息，quick_comparison-快速对比。 取值：quick_comparison
	CompareMode *string `json:"compare_mode,omitempty"`

	// 数据级对比的对象。
	CompareObject *[]CompareObjectInfo `json:"compare_object,omitempty"`

	// 对比配置项，key-value形式。 内容对比支持以下配置项： - 对比方式配置，key：contentCompareType，value：dynamic表示动态对比，static表示静态对比。 - lob字段对比类型配置，key：lobCompare，value：ignore表示忽略，length表示长度对比。  行数对比支持以下配置项： - 对比策略配置，多表归一情况下适用，key：comparePolicy，value：normal表示正常对比，manyToOne表示多对一对比。
	Options map[string]string `json:"options,omitempty"`

	// 数据级对比的对象（Cassandra灾备专用，带token信息）。
	CompareObjectWithToken *[]CompareObjectInfoWithToken `json:"compare_object_with_token,omitempty"`

	// 对比任务线程数量，当前仅cloudDataGuard-cassandra和cloudDataGuard-gausscassandra-to-gausscassandra链路支持。
	CompareTaskNum *int32 `json:"compare_task_num,omitempty"`

	// 过滤数据信息。
	CompareTransformationInfos *[]AddDataTransformationReq `json:"compare_transformation_infos,omitempty"`

	// 抽样比例，对比类型为抽样对比时填写。
	ProportionValue *float64 `json:"proportion_value,omitempty"`
}

func (o CreateDataLevelTableCompareJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataLevelTableCompareJobReq struct{}"
	}

	return strings.Join([]string{"CreateDataLevelTableCompareJobReq", string(data)}, " ")
}
