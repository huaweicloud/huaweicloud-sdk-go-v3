package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CommonDestinationDescriptor 转储任务请求体公共部分。
type CommonDestinationDescriptor struct {

	// 转储任务的名称。  任务名称由英文字母、数字、中划线和下划线组成。长度为1～64个字符。
	TaskName string `json:"task_name"`

	// 在统一身份认证服务(IAM)中创建委托的名称，DIS需要获取IAM委托信息去访问您指定的资源。创建委托的参数设置如下： - 委托类型：云服务 - 云服务：DIS - 持续时间：永久 - “所属区域”为“全局服务”，“项目”为“对象存储服务”对应的“策略”包含“Tenant Administrator”。 如果已经创建过委托，可以使用IAM服务提供的查询委托列表接口，获取有效可用的委托名称。 取值范围：长度不超过64位，且不可配置为空。  如果有在Console控制台使用转储任务，会提示自动创建委托，自动创建的委托名称为：dis_admin_agency
	AgencyName string `json:"agency_name"`

	// 根据用户配置的时间，周期性的将数据导入OBS，若某个时间段内无数据，则此时间段不会生成打包文件。  单位：秒
	DeliverTimeInterval int32 `json:"deliver_time_interval"`

	// 偏移量。  - LATEST：最大偏移量，即获取最新的数据。 - TRIM_HORIZON：最小偏移量，即读取最早的数据。
	ConsumerStrategy *CommonDestinationDescriptorConsumerStrategy `json:"consumer_strategy,omitempty"`
}

func (o CommonDestinationDescriptor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonDestinationDescriptor struct{}"
	}

	return strings.Join([]string{"CommonDestinationDescriptor", string(data)}, " ")
}

type CommonDestinationDescriptorConsumerStrategy struct {
	value string
}

type CommonDestinationDescriptorConsumerStrategyEnum struct {
	LATEST       CommonDestinationDescriptorConsumerStrategy
	TRIM_HORIZON CommonDestinationDescriptorConsumerStrategy
}

func GetCommonDestinationDescriptorConsumerStrategyEnum() CommonDestinationDescriptorConsumerStrategyEnum {
	return CommonDestinationDescriptorConsumerStrategyEnum{
		LATEST: CommonDestinationDescriptorConsumerStrategy{
			value: "LATEST",
		},
		TRIM_HORIZON: CommonDestinationDescriptorConsumerStrategy{
			value: "TRIM_HORIZON",
		},
	}
}

func (c CommonDestinationDescriptorConsumerStrategy) Value() string {
	return c.value
}

func (c CommonDestinationDescriptorConsumerStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CommonDestinationDescriptorConsumerStrategy) UnmarshalJSON(b []byte) error {
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
