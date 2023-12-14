package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ObsDestinationDescriptorRequest 转储目的地为OBS的参数列表。
type ObsDestinationDescriptorRequest struct {

	// 转储任务的名称。  任务名称由英文字母、数字、中划线和下划线组成。长度为1～64个字符。
	TaskName string `json:"task_name"`

	// 在统一身份认证服务(IAM)中创建委托的名称，DIS需要获取IAM委托信息去访问您指定的资源。创建委托的参数设置如下： - 委托类型：云服务 - 云服务：DIS - 持续时间：永久 - “所属区域”为“全局服务”，“项目”为“对象存储服务”对应的“策略”包含“Tenant Administrator”。 如果已经创建过委托，可以使用IAM服务提供的查询委托列表接口，获取有效可用的委托名称。 取值范围：长度不超过64位，且不可配置为空。  如果有在Console控制台使用转储任务，会提示自动创建委托，自动创建的委托名称为：dis_admin_agency
	AgencyName string `json:"agency_name"`

	// 根据用户配置的时间，周期性的将数据导入OBS，若某个时间段内无数据，则此时间段不会生成打包文件。  单位：秒
	DeliverTimeInterval int32 `json:"deliver_time_interval"`

	// 偏移量。  - LATEST：最大偏移量，即获取最新的数据。 - TRIM_HORIZON：最小偏移量，即读取最早的数据。
	ConsumerStrategy *ObsDestinationDescriptorRequestConsumerStrategy `json:"consumer_strategy,omitempty"`

	// 在OBS中存储通道文件的自定义目录，多级目录可用“/”进行分隔，不可以“/”开头。  取值范围：英文字母、数字、下划线和斜杠，最大长度为50个字符。  默认配置为空。
	FilePrefix *string `json:"file_prefix,omitempty"`

	// 将转储文件的生成时间使用“yyyy/MM/dd/HH/mm”格式生成分区字符串，用来定义写到OBS的Object文件所在的目录层次结构。  - N/A：置空，不使用日期时间目录。 - yyyy：年 - yyyy/MM：年/ - yyyy/MM/dd：年/月/日 - yyyy/MM/dd/HH：年/月/日/时 - yyyy/MM/dd/HH/mm：年/月/日/时/分  例如：2017/11/10/14/49，目录结构就是“2017 > 11 > 10 > 14 > 49”，“2017”表示最外层文件夹。  默认值：空  说明：  数据转储成功后，存储的目录结构为“obs_bucket_path/file_prefix/partition_format”。
	PartitionFormat *ObsDestinationDescriptorRequestPartitionFormat `json:"partition_format,omitempty"`

	// 存储该通道数据的OBS桶名称。
	ObsBucketPath string `json:"obs_bucket_path"`

	// 转储文件格式。  - text：转储目标格式为TEXT，缺省值 - parquet：转储目标格式为Parquet - carbon：转储目标格式为Carbon  说明：  “源数据类型”为“JSON”，“转储服务类型”为“OBS”时才可选择“parquet”或“carbon”格式。
	DestinationFileType *ObsDestinationDescriptorRequestDestinationFileType `json:"destination_file_type,omitempty"`

	ProcessingSchema *ProcessingSchema `json:"processing_schema,omitempty"`

	// 转储文件的记录分隔符，用于分隔写入转储文件的用户数据。  取值范围：  - 逗号 \",\"，默认值 - 分号 \";\" - 竖线 \"|\" - 换行符 \"\\n\"
	RecordDelimiter *string `json:"record_delimiter,omitempty"`
}

func (o ObsDestinationDescriptorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsDestinationDescriptorRequest struct{}"
	}

	return strings.Join([]string{"ObsDestinationDescriptorRequest", string(data)}, " ")
}

type ObsDestinationDescriptorRequestConsumerStrategy struct {
	value string
}

type ObsDestinationDescriptorRequestConsumerStrategyEnum struct {
	LATEST       ObsDestinationDescriptorRequestConsumerStrategy
	TRIM_HORIZON ObsDestinationDescriptorRequestConsumerStrategy
}

func GetObsDestinationDescriptorRequestConsumerStrategyEnum() ObsDestinationDescriptorRequestConsumerStrategyEnum {
	return ObsDestinationDescriptorRequestConsumerStrategyEnum{
		LATEST: ObsDestinationDescriptorRequestConsumerStrategy{
			value: "LATEST",
		},
		TRIM_HORIZON: ObsDestinationDescriptorRequestConsumerStrategy{
			value: "TRIM_HORIZON",
		},
	}
}

func (c ObsDestinationDescriptorRequestConsumerStrategy) Value() string {
	return c.value
}

func (c ObsDestinationDescriptorRequestConsumerStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObsDestinationDescriptorRequestConsumerStrategy) UnmarshalJSON(b []byte) error {
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

type ObsDestinationDescriptorRequestPartitionFormat struct {
	value string
}

type ObsDestinationDescriptorRequestPartitionFormatEnum struct {
	YYYY             ObsDestinationDescriptorRequestPartitionFormat
	YYYY_MM          ObsDestinationDescriptorRequestPartitionFormat
	YYYY_MM_DD       ObsDestinationDescriptorRequestPartitionFormat
	YYYY_MM_DD_HH    ObsDestinationDescriptorRequestPartitionFormat
	YYYY_MM_DD_HH_MM ObsDestinationDescriptorRequestPartitionFormat
}

func GetObsDestinationDescriptorRequestPartitionFormatEnum() ObsDestinationDescriptorRequestPartitionFormatEnum {
	return ObsDestinationDescriptorRequestPartitionFormatEnum{
		YYYY: ObsDestinationDescriptorRequestPartitionFormat{
			value: "yyyy",
		},
		YYYY_MM: ObsDestinationDescriptorRequestPartitionFormat{
			value: "yyyy/MM",
		},
		YYYY_MM_DD: ObsDestinationDescriptorRequestPartitionFormat{
			value: "yyyy/MM/dd",
		},
		YYYY_MM_DD_HH: ObsDestinationDescriptorRequestPartitionFormat{
			value: "yyyy/MM/dd/HH",
		},
		YYYY_MM_DD_HH_MM: ObsDestinationDescriptorRequestPartitionFormat{
			value: "yyyy/MM/dd/HH/mm",
		},
	}
}

func (c ObsDestinationDescriptorRequestPartitionFormat) Value() string {
	return c.value
}

func (c ObsDestinationDescriptorRequestPartitionFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObsDestinationDescriptorRequestPartitionFormat) UnmarshalJSON(b []byte) error {
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

type ObsDestinationDescriptorRequestDestinationFileType struct {
	value string
}

type ObsDestinationDescriptorRequestDestinationFileTypeEnum struct {
	TEXT    ObsDestinationDescriptorRequestDestinationFileType
	PARQUET ObsDestinationDescriptorRequestDestinationFileType
	CARBON  ObsDestinationDescriptorRequestDestinationFileType
}

func GetObsDestinationDescriptorRequestDestinationFileTypeEnum() ObsDestinationDescriptorRequestDestinationFileTypeEnum {
	return ObsDestinationDescriptorRequestDestinationFileTypeEnum{
		TEXT: ObsDestinationDescriptorRequestDestinationFileType{
			value: "text",
		},
		PARQUET: ObsDestinationDescriptorRequestDestinationFileType{
			value: "parquet",
		},
		CARBON: ObsDestinationDescriptorRequestDestinationFileType{
			value: "carbon",
		},
	}
}

func (c ObsDestinationDescriptorRequestDestinationFileType) Value() string {
	return c.value
}

func (c ObsDestinationDescriptorRequestDestinationFileType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ObsDestinationDescriptorRequestDestinationFileType) UnmarshalJSON(b []byte) error {
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
