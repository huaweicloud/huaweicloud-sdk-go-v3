package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MrsDestinationDescriptorRequest 转储目的地为MRS的参数列表。
type MrsDestinationDescriptorRequest struct {

	// 转储任务的名称。  任务名称由英文字母、数字、中划线和下划线组成。长度为1～64个字符。
	TaskName string `json:"task_name"`

	// 在统一身份认证服务(IAM)中创建委托的名称，DIS需要获取IAM委托信息去访问您指定的资源。创建委托的参数设置如下： - 委托类型：云服务 - 云服务：DIS - 持续时间：永久 - “所属区域”为“全局服务”，“项目”为“对象存储服务”对应的“策略”包含“Tenant Administrator”。 如果已经创建过委托，可以使用IAM服务提供的查询委托列表接口，获取有效可用的委托名称。 取值范围：长度不超过64位，且不可配置为空。  如果有在Console控制台使用转储任务，会提示自动创建委托，自动创建的委托名称为：dis_admin_agency
	AgencyName string `json:"agency_name"`

	// 根据用户配置的时间，周期性的将数据导入OBS，若某个时间段内无数据，则此时间段不会生成打包文件。  单位：秒
	DeliverTimeInterval int32 `json:"deliver_time_interval"`

	// 偏移量。  - LATEST：最大偏移量，即获取最新的数据。 - TRIM_HORIZON：最小偏移量，即读取最早的数据。
	ConsumerStrategy *MrsDestinationDescriptorRequestConsumerStrategy `json:"consumer_strategy,omitempty"`

	// 存储该通道数据的MRS集群名称。  说明：  仅支持非Kerberos认证的MRS集群。
	MrsClusterName string `json:"mrs_cluster_name"`

	// 存储该通道数据的MRS集群ID。
	MrsClusterId string `json:"mrs_cluster_id"`

	// 存储该通道数据的MRS集群的HDFS路径。
	MrsHdfsPath string `json:"mrs_hdfs_path"`

	// 临时存储该通道数据的OBS桶下的自定义目录，多级目录可用“/”进行分隔，不可以“/”开头。  取值范围：英文字母、数字、下划线和斜杠，最大长度为50个字符。  默认配置为空。
	FilePrefix *string `json:"file_prefix,omitempty"`

	// 在MRS集群HDFS中存储通道文件的自定义目录，多级目录可用\"/\"进行分隔。  取值范围：0~50个字符。  默认配置为空。
	HdfsPrefixFolder *string `json:"hdfs_prefix_folder,omitempty"`

	// 临时存储该通道数据的OBS桶名称。
	ObsBucketPath string `json:"obs_bucket_path"`

	// 用户数据转储失败的失效重试时间。重试时间超过该配置项配置的值，则将转储失败的数据备份至“OBS桶/ file_prefix/mrs_error”目录下。  取值范围：0~7200。  单位：秒。  默认配置为1800。  配置为“0”表示DIS服务不会在转储失败时进行重试。
	RetryDuration *string `json:"retry_duration,omitempty"`
}

func (o MrsDestinationDescriptorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MrsDestinationDescriptorRequest struct{}"
	}

	return strings.Join([]string{"MrsDestinationDescriptorRequest", string(data)}, " ")
}

type MrsDestinationDescriptorRequestConsumerStrategy struct {
	value string
}

type MrsDestinationDescriptorRequestConsumerStrategyEnum struct {
	LATEST       MrsDestinationDescriptorRequestConsumerStrategy
	TRIM_HORIZON MrsDestinationDescriptorRequestConsumerStrategy
}

func GetMrsDestinationDescriptorRequestConsumerStrategyEnum() MrsDestinationDescriptorRequestConsumerStrategyEnum {
	return MrsDestinationDescriptorRequestConsumerStrategyEnum{
		LATEST: MrsDestinationDescriptorRequestConsumerStrategy{
			value: "LATEST",
		},
		TRIM_HORIZON: MrsDestinationDescriptorRequestConsumerStrategy{
			value: "TRIM_HORIZON",
		},
	}
}

func (c MrsDestinationDescriptorRequestConsumerStrategy) Value() string {
	return c.value
}

func (c MrsDestinationDescriptorRequestConsumerStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MrsDestinationDescriptorRequestConsumerStrategy) UnmarshalJSON(b []byte) error {
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
