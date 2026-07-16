package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type IntranetConnectionRequest struct {

	// **参数解释：** 服务ID，在[创建服务](CreateInferService.xml)时即可在返回体中获取，也可通过[查询服务列表](ListInferServices.xml)获取当前用户拥有的服务，其中service_id字段即为服务ID。 **约束限制：** type为SERVICE时，必填。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ServiceId *string `json:"service_id,omitempty"`

	// **参数解释：** 内网访问场景。 **约束限制：** 不涉及。 **取值范围：** - VPC：用户VPC网络接入场景 - POOL：用户资源池网络接入场景 **默认取值：** 不涉及。
	Scene string `json:"scene"`

	// **参数解释：** VPC ID，VPC场景需要填写。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释：** 子网 ID，VPC场景需要填写。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释：** 资源池id POOL场景需要填写。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释：** 内网访问接入粒度，不填默认为SERVICE **约束限制：** 不涉及。 **取值范围：** - SERVICE：服务粒度。 - GLOBAL：global粒度。 **默认取值：** 默认为SERVICE。
	Type *IntranetConnectionRequestType `json:"type,omitempty"`

	// **参数解释：** 服务绑定的dispatcher组ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	DispatcherGroupId *string `json:"dispatcher_group_id,omitempty"`

	// **参数解释：** 自定义URL，格式为：{协议}://{域名}{路径} **约束限制：** url个数不超过10个，单个url长度不超过1024。 **取值范围：** - 协议范围：http，https，wss，ws。 - 域名范围：支持域名或IP:端口。域名长度不超过63，包含字母、数字、中划线（-)且不能以中划线（-)开头或结尾，顶级域名不能包含数字；端口范围为1-65535。 - 路径范围：斜杠（/）开头，仅包含字母、数字、点号（.）、中划线（-)、下划线（_）、斜杠（/）的路径。 **默认取值：** 不涉及。
	CustomUrls *[]string `json:"custom_urls,omitempty"`
}

func (o IntranetConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntranetConnectionRequest struct{}"
	}

	return strings.Join([]string{"IntranetConnectionRequest", string(data)}, " ")
}

type IntranetConnectionRequestType struct {
	value string
}

type IntranetConnectionRequestTypeEnum struct {
	SERVICE IntranetConnectionRequestType
	GLOBAL  IntranetConnectionRequestType
}

func GetIntranetConnectionRequestTypeEnum() IntranetConnectionRequestTypeEnum {
	return IntranetConnectionRequestTypeEnum{
		SERVICE: IntranetConnectionRequestType{
			value: "SERVICE",
		},
		GLOBAL: IntranetConnectionRequestType{
			value: "GLOBAL",
		},
	}
}

func (c IntranetConnectionRequestType) Value() string {
	return c.value
}

func (c IntranetConnectionRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IntranetConnectionRequestType) UnmarshalJSON(b []byte) error {
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
