package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateAccessConfigRequestBody 创建日志接入请求体
type CreateAccessConfigRequestBody struct {

	// 日志接入名称。 满足正则表达式：^(?!\\.)(?!_)(?!.*?\\.$)[\\u4e00-\\u9fa5a-zA-Z0-9-_.]{1,64}$
	AccessConfigName string `json:"access_config_name"`

	// 日志接入类型。AGENT：ECS接入,K8S_CCE:CCE接入
	AccessConfigType CreateAccessConfigRequestBodyAccessConfigType `json:"access_config_type"`

	AccessConfigDetail *AccessConfigDeatilCreate `json:"access_config_detail"`

	LogInfo *AccessConfigBaseLogInfoCreate `json:"log_info"`

	HostGroupInfo *AccessConfigHostGroupIdListCreate `json:"host_group_info,omitempty"`

	// 标签信息。KEY不能重复,最多20个标签
	AccessConfigTag *[]AccessConfigTag `json:"access_config_tag,omitempty"`

	// 二进制采集
	BinaryCollect *bool `json:"binary_collect,omitempty"`

	// 日志拆分
	LogSplit *bool `json:"log_split,omitempty"`

	// CCE集群ID，当CCE类型时，为必填
	ClusterId *string `json:"cluster_id,omitempty"`

	// 是否增量采集 true 为是   false为否（全量采集）
	IncrementalCollect *bool `json:"incremental_collect,omitempty"`

	// 编码格式，支持UTF-8，GDB默认UTF-8
	EncodingFormat *string `json:"encoding_format,omitempty"`

	// IC结构化解析类型包括 ：SINGLE_LINE 单行全文，MULTI_LINE 多行全文，REGEX 单行正则，MULTI_REGEX 多行正则，SPLIT 分隔符，JSON JSON解析，NGINX nginx解析， COMPOSE组合解析
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

	// 日志接入自建软件来源
	AccessConfigTypeSource *string `json:"access_config_type_source,omitempty"`
}

func (o CreateAccessConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessConfigRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAccessConfigRequestBody", string(data)}, " ")
}

type CreateAccessConfigRequestBodyAccessConfigType struct {
	value string
}

type CreateAccessConfigRequestBodyAccessConfigTypeEnum struct {
	AGENT    CreateAccessConfigRequestBodyAccessConfigType
	K8_S_CCE CreateAccessConfigRequestBodyAccessConfigType
}

func GetCreateAccessConfigRequestBodyAccessConfigTypeEnum() CreateAccessConfigRequestBodyAccessConfigTypeEnum {
	return CreateAccessConfigRequestBodyAccessConfigTypeEnum{
		AGENT: CreateAccessConfigRequestBodyAccessConfigType{
			value: "AGENT",
		},
		K8_S_CCE: CreateAccessConfigRequestBodyAccessConfigType{
			value: "K8S_CCE",
		},
	}
}

func (c CreateAccessConfigRequestBodyAccessConfigType) Value() string {
	return c.value
}

func (c CreateAccessConfigRequestBodyAccessConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAccessConfigRequestBodyAccessConfigType) UnmarshalJSON(b []byte) error {
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
