package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePublicipOption 弹性公网IP对象
type UpdatePublicipOption struct {

	// 功能说明：端口id  约束：必须是存在的端口id，如果不带该参数或者值为空时为解除绑定弹性公网IP，如果该端口不存在或端口已绑定弹性公网IP则会提示出错。  和ip_version字段互斥，不能同时更新。如果alias和port_id都携带，只有alias生效。
	PortId *string `json:"port_id,omitempty"`

	// 功能说明：IP版本信息  取值范围：4和6  4：IPv4  6：IPv6  约束：必须是系统支持的IP版本类型，和port_id互斥，不能同时更新。
	IpVersion *UpdatePublicipOptionIpVersion `json:"ip_version,omitempty"`

	// 功能说明：弹性公网IP名称 取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点） 约束：如果alias和port_id都携带，只有alias生效。
	Alias *string `json:"alias,omitempty"`
}

func (o UpdatePublicipOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicipOption struct{}"
	}

	return strings.Join([]string{"UpdatePublicipOption", string(data)}, " ")
}

type UpdatePublicipOptionIpVersion struct {
	value int32
}

type UpdatePublicipOptionIpVersionEnum struct {
	E_4 UpdatePublicipOptionIpVersion
	E_6 UpdatePublicipOptionIpVersion
}

func GetUpdatePublicipOptionIpVersionEnum() UpdatePublicipOptionIpVersionEnum {
	return UpdatePublicipOptionIpVersionEnum{
		E_4: UpdatePublicipOptionIpVersion{
			value: 4,
		}, E_6: UpdatePublicipOptionIpVersion{
			value: 6,
		},
	}
}

func (c UpdatePublicipOptionIpVersion) Value() int32 {
	return c.value
}

func (c UpdatePublicipOptionIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePublicipOptionIpVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
