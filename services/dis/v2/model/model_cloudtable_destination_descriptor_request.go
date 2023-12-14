package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CloudtableDestinationDescriptorRequest 转储目的地为CloudTable的参数列表。
type CloudtableDestinationDescriptorRequest struct {

	// 转储任务的名称。  任务名称由英文字母、数字、中划线和下划线组成。长度为1～64个字符。
	TaskName string `json:"task_name"`

	// 在统一身份认证服务(IAM)中创建委托的名称，DIS需要获取IAM委托信息去访问您指定的资源。创建委托的参数设置如下： - 委托类型：云服务 - 云服务：DIS - 持续时间：永久 - “所属区域”为“全局服务”，“项目”为“对象存储服务”对应的“策略”包含“Tenant Administrator”。 如果已经创建过委托，可以使用IAM服务提供的查询委托列表接口，获取有效可用的委托名称。 取值范围：长度不超过64位，且不可配置为空。  如果有在Console控制台使用转储任务，会提示自动创建委托，自动创建的委托名称为：dis_admin_agency
	AgencyName string `json:"agency_name"`

	// 根据用户配置的时间，周期性的将数据导入OBS，若某个时间段内无数据，则此时间段不会生成打包文件。  单位：秒
	DeliverTimeInterval int32 `json:"deliver_time_interval"`

	// 偏移量。  - LATEST：最大偏移量，即获取最新的数据。 - TRIM_HORIZON：最小偏移量，即读取最早的数据。
	ConsumerStrategy *CloudtableDestinationDescriptorRequestConsumerStrategy `json:"consumer_strategy,omitempty"`

	// 存储该通道数据的CloudTable集群名称。  如果选择转储OpenTSDB，则集群必须开启OpenTSDB。
	CloudtableClusterName string `json:"cloudtable_cluster_name"`

	// 存储该通道数据的CloudTable集群ID。  如果选择转储OpenTSDB，则集群必须开启OpenTSDB。
	CloudtableClusterId string `json:"cloudtable_cluster_id"`

	// 转储HBase时必选，表示存储该通道数据的CloudTable集群HBase表名称。
	CloudtableTableName *string `json:"cloudtable_table_name,omitempty"`

	CloudtableSchema *CloudtableSchema `json:"cloudtable_schema,omitempty"`

	// 转储OpenTSDB时必选，与“cloudtable_schema”二选一，表示CloudTable集群OpenTSDB数据的Schema配置。用于将通道内的JSON数据进行格式转换并导入Cloudtable的OpenTSDB。
	OpentsdbSchema *[]OpenTsdbSchema `json:"opentsdb_schema,omitempty"`

	// 转储HBase的rowkey分隔符，用于分隔生成rowKey的用户数据。取值范围：”, ”、 ”. ”、 ”|”、 ”; ”、 ”\\”、 ”-”、 ”_”、 ”~”  缺省值：”.”
	CloudtableRowKeyDelimiter *string `json:"cloudtable_row_key_delimiter,omitempty"`

	// 用户数据转储CloudTable服务失败时，可选择将转储失败的数据备份至OBS服务，此参数为OBS服务的桶名称。
	ObsBackupBucketPath *string `json:"obs_backup_bucket_path,omitempty"`

	// 用户数据转储CloudTable服务失败时，可选择将转储失败的数据备份至OBS服务，此参数为OBS桶下的自定义目录，多级目录可用“/”进行分隔，不可以“/”开头。  取值范围：英文字母、数字和下划线。  最大长度：最大长度为50个字符。  默认配置为空。
	BackupFilePrefix *string `json:"backup_file_prefix,omitempty"`

	// 用户数据导入CloudTable服务失败的失效重试时间。超出此时效，转储CloudTable失败的数据将备份至“OBS桶/ backup_file_prefix /cloudtable_error”或“OBS桶/ backup_file_prefix /opentsdb_error”目录下。  取值范围： 0～7200。  单位：秒。  默认配置为1800。
	RetryDuration *string `json:"retry_duration,omitempty"`
}

func (o CloudtableDestinationDescriptorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudtableDestinationDescriptorRequest struct{}"
	}

	return strings.Join([]string{"CloudtableDestinationDescriptorRequest", string(data)}, " ")
}

type CloudtableDestinationDescriptorRequestConsumerStrategy struct {
	value string
}

type CloudtableDestinationDescriptorRequestConsumerStrategyEnum struct {
	LATEST       CloudtableDestinationDescriptorRequestConsumerStrategy
	TRIM_HORIZON CloudtableDestinationDescriptorRequestConsumerStrategy
}

func GetCloudtableDestinationDescriptorRequestConsumerStrategyEnum() CloudtableDestinationDescriptorRequestConsumerStrategyEnum {
	return CloudtableDestinationDescriptorRequestConsumerStrategyEnum{
		LATEST: CloudtableDestinationDescriptorRequestConsumerStrategy{
			value: "LATEST",
		},
		TRIM_HORIZON: CloudtableDestinationDescriptorRequestConsumerStrategy{
			value: "TRIM_HORIZON",
		},
	}
}

func (c CloudtableDestinationDescriptorRequestConsumerStrategy) Value() string {
	return c.value
}

func (c CloudtableDestinationDescriptorRequestConsumerStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudtableDestinationDescriptorRequestConsumerStrategy) UnmarshalJSON(b []byte) error {
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
