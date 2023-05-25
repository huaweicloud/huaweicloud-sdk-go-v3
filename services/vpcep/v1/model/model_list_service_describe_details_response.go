package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ListServiceDescribeDetailsResponse struct {

	// 终端节点服务的ID，唯一标识。
	Id *string `json:"id,omitempty"`

	// 终端节点服务的名称。
	ServiceName *string `json:"service_name,omitempty"`

	// 终端节点服务类型。仅支持将用户私有服务创建为interface类型的终端节点服务。 ● gataway：由运维人员配置。用户无需创建，可直接使用。 ● interface：包括运维人员配置的云服务和用户自己创建的私有服务。 其中，运维人员配置的云服务无需创建，用户可直接使用。 您可以通过创建终端节点创建访问Gateway和Interface类型终端节点服务的终端节点。
	ServiceType *ListServiceDescribeDetailsResponseServiceType `json:"service_type,omitempty"`

	// 终端节点服务的创建时间。 采用UTC时间格式，格式为：YYYY-MMDDTHH:MM:SSZ
	CreatedAt *string `json:"created_at,omitempty"`

	// 连接该终端节点服务的终端节点是否计费。 ● true：计费 ● false：不计费
	IsCharge *bool `json:"is_charge,omitempty"`

	// 终端节点对应Pool的Public Border Group信息
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 是否开启终端节点策略。 ● false：不支持设置终端节点策略 ● true：支持设置终端节点策略 默认为false 是否开启终端节点策略。 ● false：不支持设置终端节点策略 ● true：支持设置终端节点策略 默认为false
	EnablePolicy   *bool `json:"enable_policy,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListServiceDescribeDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceDescribeDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListServiceDescribeDetailsResponse", string(data)}, " ")
}

type ListServiceDescribeDetailsResponseServiceType struct {
	value string
}

type ListServiceDescribeDetailsResponseServiceTypeEnum struct {
	INTERFACE ListServiceDescribeDetailsResponseServiceType
}

func GetListServiceDescribeDetailsResponseServiceTypeEnum() ListServiceDescribeDetailsResponseServiceTypeEnum {
	return ListServiceDescribeDetailsResponseServiceTypeEnum{
		INTERFACE: ListServiceDescribeDetailsResponseServiceType{
			value: "interface",
		},
	}
}

func (c ListServiceDescribeDetailsResponseServiceType) Value() string {
	return c.value
}

func (c ListServiceDescribeDetailsResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListServiceDescribeDetailsResponseServiceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
