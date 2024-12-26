package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessConfigRequestBody 修改日志接入请求体
type UpdateAccessConfigRequestBody struct {

	// 日志接入ID
	AccessConfigId string `json:"access_config_id"`

	// 日志接入名称。 满足正则表达式：^(?!.)(?!)(?!.*?.$)[\\u4e00-\\u9fa5a-zA-Z0-9-.]{1,64}$
	AccessConfigName *string `json:"access_config_name,omitempty"`

	AccessConfigDetail *AccessConfigDeatilUpdate `json:"access_config_detail,omitempty"`

	HostGroupInfo *AccessConfigHostGroupIdList `json:"host_group_info,omitempty"`

	// 标签信息。KEY不能重复,最多20个标签
	AccessConfigTag *[]AccessConfigTag `json:"access_config_tag,omitempty"`

	// 日志拆分
	LogSplit *bool `json:"log_split,omitempty"`

	// 二进制采集
	BinaryCollect *bool `json:"binary_collect,omitempty"`

	// CCE集群ID，CCE类型时，为必填
	ClusterId *string `json:"cluster_id,omitempty"`

	// 是否增量采集 true为是 false为否（全量采集）
	IncrementalCollect *bool `json:"incremental_collect,omitempty"`

	// 编码格式，默认UTF-8
	EncodingFormat *string `json:"encoding_format,omitempty"`

	// IC结构化解析类型
	ProcessorType *string `json:"processor_type,omitempty"`

	// 示例日志
	DemoLog *string `json:"demo_log,omitempty"`

	// 示例日志解析字段
	DemoFields *[]DemoFieldAccess `json:"demo_fields,omitempty"`

	// IC结构化解析器
	Processors *[]Processor `json:"processors,omitempty"`

	// ServiceStage应用ID
	ApplicationId *string `json:"application_id,omitempty"`

	// ServiceStage环境ID
	EnvironmentId *string `json:"environment_id,omitempty"`

	// ServiceStage组件ID
	ComponentId *[]string `json:"component_id,omitempty"`
}

func (o UpdateAccessConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAccessConfigRequestBody", string(data)}, " ")
}
