package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateAccessConfigResponse Response Object
type UpdateAccessConfigResponse struct {

	// 日志接入ID
	AccessConfigId *string `json:"access_config_id,omitempty"`

	// 日志接入名称
	AccessConfigName *string `json:"access_config_name,omitempty"`

	// 日志接入类型。AGENT：ECS接入  K8S_CCE: CCE接入
	AccessConfigType *UpdateAccessConfigResponseAccessConfigType `json:"access_config_type,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	AccessConfigDetail *AccessConfigDeatilCreate `json:"access_config_detail,omitempty"`

	LogInfo *AccessConfigQueryLogInfo `json:"log_info,omitempty"`

	HostGroupInfo *AccessConfigHostGroupIdList `json:"host_group_info,omitempty"`

	// 标签信息。KEY不能重复,最多20个标签
	AccessConfigTag *[]AccessConfigTag `json:"access_config_tag,omitempty"`

	// 二进制采集
	LogSplit *bool `json:"log_split,omitempty"`

	// 日志拆分
	BinaryCollect *bool `json:"binary_collect,omitempty"`

	// CCE集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 编码格式，默认UTF-8
	EncodingFormat *string `json:"encoding_format,omitempty"`

	// 采集策略：增量/全量
	IncrementalCollect *bool `json:"incremental_collect,omitempty"`

	// IC结构化解析类型
	ProcessorType *string `json:"processor_type,omitempty"`

	// 示例日志
	DemoLog *string `json:"demo_log,omitempty"`

	// 示例日志解析字段
	DemoFields *[]DemoFieldAccess `json:"demo_fields,omitempty"`

	// IC结构化解析器
	Processors *[]Processor `json:"processors,omitempty"`

	// **参数解释：** 拆分日志大小。 **取值范围：** 不涉及。
	LogSplitSize *int32 `json:"log_split_size,omitempty"`

	// **参数解释：** 采集路径递归最大深度。 **取值范围：** 不涉及。
	RecursiveDepth *int32 `json:"recursive_depth,omitempty"`

	// **参数解释：** 日志接入自建软件来源。 **取值范围：** - ECS - CCE - BMS - K8S - ServiceStageHost - ServiceStage
	AccessConfigTypeSource *string `json:"access_config_type_source,omitempty"`

	// ServiceStage应用ID
	ApplicationId *string `json:"application_id,omitempty"`

	// ServiceStage环境ID
	EnvironmentId *string `json:"environment_id,omitempty"`

	// ServiceStage组件ID
	ComponentId    *[]string `json:"component_id,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateAccessConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateAccessConfigResponse", string(data)}, " ")
}

type UpdateAccessConfigResponseAccessConfigType struct {
	value string
}

type UpdateAccessConfigResponseAccessConfigTypeEnum struct {
	AGENT    UpdateAccessConfigResponseAccessConfigType
	K8_S_CCE UpdateAccessConfigResponseAccessConfigType
}

func GetUpdateAccessConfigResponseAccessConfigTypeEnum() UpdateAccessConfigResponseAccessConfigTypeEnum {
	return UpdateAccessConfigResponseAccessConfigTypeEnum{
		AGENT: UpdateAccessConfigResponseAccessConfigType{
			value: "AGENT",
		},
		K8_S_CCE: UpdateAccessConfigResponseAccessConfigType{
			value: "K8S_CCE",
		},
	}
}

func (c UpdateAccessConfigResponseAccessConfigType) Value() string {
	return c.value
}

func (c UpdateAccessConfigResponseAccessConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAccessConfigResponseAccessConfigType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
