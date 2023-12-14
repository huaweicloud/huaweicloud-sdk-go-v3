package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DwsDestinationDescriptorRequest 转储目的地为DWS的参数列表。
type DwsDestinationDescriptorRequest struct {

	// 转储任务的名称。  任务名称由英文字母、数字、中划线和下划线组成。长度为1～64个字符。
	TaskName string `json:"task_name"`

	// 在统一身份认证服务(IAM)中创建委托的名称，DIS需要获取IAM委托信息去访问您指定的资源。创建委托的参数设置如下： - 委托类型：云服务 - 云服务：DIS - 持续时间：永久 - “所属区域”为“全局服务”，“项目”为“对象存储服务”对应的“策略”包含“Tenant Administrator”。 如果已经创建过委托，可以使用IAM服务提供的查询委托列表接口，获取有效可用的委托名称。 取值范围：长度不超过64位，且不可配置为空。  如果有在Console控制台使用转储任务，会提示自动创建委托，自动创建的委托名称为：dis_admin_agency
	AgencyName string `json:"agency_name"`

	// 根据用户配置的时间，周期性的将数据导入OBS，若某个时间段内无数据，则此时间段不会生成打包文件。  单位：秒
	DeliverTimeInterval int32 `json:"deliver_time_interval"`

	// 偏移量。  - LATEST：最大偏移量，即获取最新的数据。 - TRIM_HORIZON：最小偏移量，即读取最早的数据。
	ConsumerStrategy *DwsDestinationDescriptorRequestConsumerStrategy `json:"consumer_strategy,omitempty"`

	// 存储该通道数据的DWS集群名称。
	DwsClusterName string `json:"dws_cluster_name"`

	// 存储该通道数据的DWS集群ID。
	DwsClusterId string `json:"dws_cluster_id"`

	// 存储该通道数据的DWS数据库名称。
	DwsDatabaseName string `json:"dws_database_name"`

	// 存储该通道数据的DWS数据库模式。
	DwsSchema string `json:"dws_schema"`

	// 存储该通道数据的DWS数据库模式下的数据表。
	DwsTableName string `json:"dws_table_name"`

	// 用户数据的字段分隔符，根据此分隔符分隔用户数据插入DWS数据表的相应列。  取值范围：“，”、“；”和“|”三种字符中的一个。
	DwsDelimiter string `json:"dws_delimiter"`

	// 存储该通道数据的DWS数据库的用户名。
	UserName string `json:"user_name"`

	// 存储该通道数据的DWS数据库的密码。
	UserPassword string `json:"user_password"`

	// 用户在密钥管理服务（简称KMS）创建的用户主密钥名称，用于加密存储DWS数据库的密码。
	KmsUserKeyName string `json:"kms_user_key_name"`

	// 用户在密钥管理服务（简称KMS）创建的用户主密钥ID，用于加密存储DWS数据库的密码。
	KmsUserKeyId string `json:"kms_user_key_id"`

	// 临时存储该通道数据的OBS桶名称。
	ObsBucketPath string `json:"obs_bucket_path"`

	// 临时存储该通道数据的OBS桶下的自定义目录，多级目录可用“/”进行分隔，不可以“/”开头。  取值范围：英文字母、数字、下划线和斜杠，最大长度为50个字符。  默认配置为空。
	FilePrefix *string `json:"file_prefix,omitempty"`

	// 用户数据导入DWS集群失败的重试失效时间。超出此配置项配置的时间，转储DWS失败的数据将备份至“OBS桶/ file_prefix/dws_error”目录下。  取值范围： 0～7200。  单位：秒。  默认配置为1800。
	RetryDuration *string `json:"retry_duration,omitempty"`

	// 指定要转储到DWS表中的列，为null或者为空则默认全列。比如“c1,c2”表示将Schema中c1和c2这两列转储到DWS。  默认为空。
	DwsTableColumns *string `json:"dws_table_columns,omitempty"`

	Options *Options `json:"options,omitempty"`
}

func (o DwsDestinationDescriptorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DwsDestinationDescriptorRequest struct{}"
	}

	return strings.Join([]string{"DwsDestinationDescriptorRequest", string(data)}, " ")
}

type DwsDestinationDescriptorRequestConsumerStrategy struct {
	value string
}

type DwsDestinationDescriptorRequestConsumerStrategyEnum struct {
	LATEST       DwsDestinationDescriptorRequestConsumerStrategy
	TRIM_HORIZON DwsDestinationDescriptorRequestConsumerStrategy
}

func GetDwsDestinationDescriptorRequestConsumerStrategyEnum() DwsDestinationDescriptorRequestConsumerStrategyEnum {
	return DwsDestinationDescriptorRequestConsumerStrategyEnum{
		LATEST: DwsDestinationDescriptorRequestConsumerStrategy{
			value: "LATEST",
		},
		TRIM_HORIZON: DwsDestinationDescriptorRequestConsumerStrategy{
			value: "TRIM_HORIZON",
		},
	}
}

func (c DwsDestinationDescriptorRequestConsumerStrategy) Value() string {
	return c.value
}

func (c DwsDestinationDescriptorRequestConsumerStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DwsDestinationDescriptorRequestConsumerStrategy) UnmarshalJSON(b []byte) error {
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
