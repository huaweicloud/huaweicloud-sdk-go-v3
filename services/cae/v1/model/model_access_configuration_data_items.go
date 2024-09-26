package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessConfigurationDataItems 访问方式配置项。
type AccessConfigurationDataItems struct {

	// 配置模式。 - 如果operator值为空，则表示使用全量覆盖模式进行配置，否则表示使用增删改模式进行配置。且此级列表的所有元素的operator值必须同时全为空或者非空。 - 当使用增删改模式时，operator取值支持\"add\",\"copy\",\"modify\",\"delete\"，分别表示新增，复制指定uid的元素修改后新增，修改指定uid的元素，删除指定uid的元素。 - 当operator取值为\"copy\",\"modify\",\"delete\"时，uid的值必须为非空，且存在于最后一次生效的配置中。 - 当operator取值为\"copy\",\"modify\"时，与operator同级别的字段中除uid外的所有字段如不写，置空或者为空列表，则表示保留在最后一次生效配置中指定uid的元素的同一字段的值。
	Operator *string `json:"operator,omitempty"`

	// 访问方式的uid。
	Uid *string `json:"uid,omitempty"`

	Metadata *AccessConfigurationMetadata `json:"metadata,omitempty"`

	// 访问方式类型。
	Type *AccessConfigurationDataItemsType `json:"type,omitempty"`

	// 内网访问方式域名。
	DomainNames *[]string `json:"domain_names,omitempty"`

	AccessControl *AccessControl `json:"access_control,omitempty"`

	// 访问方式配置端口、协议、证书、URL路径等信息列表。
	Ports *[]AccessConfigurationPort `json:"ports,omitempty"`

	// 用户选择的elb的ID。
	ElbId *string `json:"elb_id,omitempty"`

	// 响应体参数，用户选择的elb的公网ip。
	PublicIp *string `json:"public_ip,omitempty"`

	// 响应体参数，用户选择的elb的私网ip。
	PrivateIp *string `json:"private_ip,omitempty"`
}

func (o AccessConfigurationDataItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigurationDataItems struct{}"
	}

	return strings.Join([]string{"AccessConfigurationDataItems", string(data)}, " ")
}

type AccessConfigurationDataItemsType struct {
	value string
}

type AccessConfigurationDataItemsTypeEnum struct {
	CLUSTER_IP    AccessConfigurationDataItemsType
	LOAD_BALANCER AccessConfigurationDataItemsType
	INGRESS       AccessConfigurationDataItemsType
}

func GetAccessConfigurationDataItemsTypeEnum() AccessConfigurationDataItemsTypeEnum {
	return AccessConfigurationDataItemsTypeEnum{
		CLUSTER_IP: AccessConfigurationDataItemsType{
			value: "ClusterIP",
		},
		LOAD_BALANCER: AccessConfigurationDataItemsType{
			value: "LoadBalancer",
		},
		INGRESS: AccessConfigurationDataItemsType{
			value: "Ingress",
		},
	}
}

func (c AccessConfigurationDataItemsType) Value() string {
	return c.value
}

func (c AccessConfigurationDataItemsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigurationDataItemsType) UnmarshalJSON(b []byte) error {
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
