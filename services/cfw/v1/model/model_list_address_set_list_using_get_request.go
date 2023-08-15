package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAddressSetListUsingGetRequest Request Object
type ListAddressSetListUsingGetRequest struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用查询防火墙实例接口获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。具体可参考APIExlorer和帮助中心FAQ。
	ObjectId string `json:"object_id"`

	// 关键字
	KeyWord *string `json:"key_word,omitempty"`

	// 每页显示个数
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// ip地址
	Address *string `json:"address,omitempty"`

	// 地址类型0 ipv4,1 ipv6
	AddressType *ListAddressSetListUsingGetRequestAddressType `json:"address_type,omitempty"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用查询防火墙实例接口获得。具体可参考APIExlorer和帮助中心FAQ。默认情况下，fw_instance_Id为空时，返回帐号下第一个墙的信息；fw_instance_Id非空时，返回与fw_instance_Id对应墙的信息。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`
}

func (o ListAddressSetListUsingGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressSetListUsingGetRequest struct{}"
	}

	return strings.Join([]string{"ListAddressSetListUsingGetRequest", string(data)}, " ")
}

type ListAddressSetListUsingGetRequestAddressType struct {
	value int32
}

type ListAddressSetListUsingGetRequestAddressTypeEnum struct {
	E_0 ListAddressSetListUsingGetRequestAddressType
	E_1 ListAddressSetListUsingGetRequestAddressType
}

func GetListAddressSetListUsingGetRequestAddressTypeEnum() ListAddressSetListUsingGetRequestAddressTypeEnum {
	return ListAddressSetListUsingGetRequestAddressTypeEnum{
		E_0: ListAddressSetListUsingGetRequestAddressType{
			value: 0,
		}, E_1: ListAddressSetListUsingGetRequestAddressType{
			value: 1,
		},
	}
}

func (c ListAddressSetListUsingGetRequestAddressType) Value() int32 {
	return c.value
}

func (c ListAddressSetListUsingGetRequestAddressType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAddressSetListUsingGetRequestAddressType) UnmarshalJSON(b []byte) error {
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
