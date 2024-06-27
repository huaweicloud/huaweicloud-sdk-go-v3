package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAddressSetsRequest Request Object
type ListAddressSetsRequest struct {

	// 互联网边界防护对象id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，type为0的为互联网边界防护对象id。
	ObjectId string `json:"object_id"`

	// 关键字
	KeyWord *string `json:"key_word,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// ip地址
	Address *string `json:"address,omitempty"`

	// 地址类型0 ipv4,1 ipv6
	AddressType *ListAddressSetsRequestAddressType `json:"address_type,omitempty"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，默认情况下，fw_instance_Id为空时，返回账号下第一个墙的信息；fw_instance_Id非空时，返回与fw_instance_Id对应墙的信息。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 查询地址组类型，0表示自定义地址组，1表示预定义地址组
	QueryAddressSetType *int32 `json:"query_address_set_type,omitempty"`

	// 地址组类型，0表示自定义地址组，1表示WAF回源IP地址组，2表示DDoS回源IP地址组，3表示NAT64转换地址组
	AddressSetType *int32 `json:"address_set_type,omitempty"`
}

func (o ListAddressSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressSetsRequest struct{}"
	}

	return strings.Join([]string{"ListAddressSetsRequest", string(data)}, " ")
}

type ListAddressSetsRequestAddressType struct {
	value int32
}

type ListAddressSetsRequestAddressTypeEnum struct {
	E_0 ListAddressSetsRequestAddressType
	E_1 ListAddressSetsRequestAddressType
}

func GetListAddressSetsRequestAddressTypeEnum() ListAddressSetsRequestAddressTypeEnum {
	return ListAddressSetsRequestAddressTypeEnum{
		E_0: ListAddressSetsRequestAddressType{
			value: 0,
		}, E_1: ListAddressSetsRequestAddressType{
			value: 1,
		},
	}
}

func (c ListAddressSetsRequestAddressType) Value() int32 {
	return c.value
}

func (c ListAddressSetsRequestAddressType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAddressSetsRequestAddressType) UnmarshalJSON(b []byte) error {
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
