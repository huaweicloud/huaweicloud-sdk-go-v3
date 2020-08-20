/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// 弹性公网IP对象
type UpdatePublicipOption struct {
	// 功能说明：端口id  约束：必须是存在的端口id，如果不带该参数或者值为空时为解除绑定弹性公网IP，如果该端口不存在或端口已绑定弹性公网IP则会提示出错。  和ip_version字段互斥，不能同时更新。
	PortId *string `json:"port_id,omitempty"`
	// 功能说明：IP版本信息  取值范围：4和6  4：IPv4  6：IPv6  约束：必须是系统支持的IP版本类型，和port_id互斥，不能同时更新。
	IpVersion UpdatePublicipOptionIpVersion `json:"ip_version,omitempty"`
}

func (o UpdatePublicipOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePublicipOption", string(data)}, " ")
}

type UpdatePublicipOptionIpVersion struct {
	value int32
}

type UpdatePublicipOptionIpVersionEnum struct {
	_4 UpdatePublicipOptionIpVersion
	_6 UpdatePublicipOptionIpVersion
}

func GetUpdatePublicipOptionIpVersionEnum() UpdatePublicipOptionIpVersionEnum {
	return UpdatePublicipOptionIpVersionEnum{
		_4: UpdatePublicipOptionIpVersion{
			value: 4,
		}, _6: UpdatePublicipOptionIpVersion{
			value: 6,
		},
	}
}

func (c UpdatePublicipOptionIpVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdatePublicipOptionIpVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
