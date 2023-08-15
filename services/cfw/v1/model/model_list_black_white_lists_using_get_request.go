package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBlackWhiteListsUsingGetRequest Request Object
type ListBlackWhiteListsUsingGetRequest struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用查询防火墙实例接口获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。具体可参考APIExlorer和帮助中心FAQ。
	ObjectId string `json:"object_id"`

	// 黑白名单类型4：黑名单，5：白名单
	ListType ListBlackWhiteListsUsingGetRequestListType `json:"list_type"`

	// IP地址类型0：ipv4,1:ipv6,2:domain
	AddressType *ListBlackWhiteListsUsingGetRequestAddressType `json:"address_type,omitempty"`

	// ip地址
	Address *string `json:"address,omitempty"`

	// 端口
	Port *string `json:"port,omitempty"`

	// 每页显示个数
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用查询防火墙实例接口获得。具体可参考APIExlorer和帮助中心FAQ。默认情况下，fw_instance_Id为空时，返回帐号下第一个墙的信息；fw_instance_Id非空时，返回与fw_instance_Id对应墙的信息。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`
}

func (o ListBlackWhiteListsUsingGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBlackWhiteListsUsingGetRequest struct{}"
	}

	return strings.Join([]string{"ListBlackWhiteListsUsingGetRequest", string(data)}, " ")
}

type ListBlackWhiteListsUsingGetRequestListType struct {
	value int32
}

type ListBlackWhiteListsUsingGetRequestListTypeEnum struct {
	E_4 ListBlackWhiteListsUsingGetRequestListType
	E_5 ListBlackWhiteListsUsingGetRequestListType
}

func GetListBlackWhiteListsUsingGetRequestListTypeEnum() ListBlackWhiteListsUsingGetRequestListTypeEnum {
	return ListBlackWhiteListsUsingGetRequestListTypeEnum{
		E_4: ListBlackWhiteListsUsingGetRequestListType{
			value: 4,
		}, E_5: ListBlackWhiteListsUsingGetRequestListType{
			value: 5,
		},
	}
}

func (c ListBlackWhiteListsUsingGetRequestListType) Value() int32 {
	return c.value
}

func (c ListBlackWhiteListsUsingGetRequestListType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBlackWhiteListsUsingGetRequestListType) UnmarshalJSON(b []byte) error {
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

type ListBlackWhiteListsUsingGetRequestAddressType struct {
	value int32
}

type ListBlackWhiteListsUsingGetRequestAddressTypeEnum struct {
	E_0 ListBlackWhiteListsUsingGetRequestAddressType
	E_1 ListBlackWhiteListsUsingGetRequestAddressType
	E_2 ListBlackWhiteListsUsingGetRequestAddressType
}

func GetListBlackWhiteListsUsingGetRequestAddressTypeEnum() ListBlackWhiteListsUsingGetRequestAddressTypeEnum {
	return ListBlackWhiteListsUsingGetRequestAddressTypeEnum{
		E_0: ListBlackWhiteListsUsingGetRequestAddressType{
			value: 0,
		}, E_1: ListBlackWhiteListsUsingGetRequestAddressType{
			value: 1,
		}, E_2: ListBlackWhiteListsUsingGetRequestAddressType{
			value: 2,
		},
	}
}

func (c ListBlackWhiteListsUsingGetRequestAddressType) Value() int32 {
	return c.value
}

func (c ListBlackWhiteListsUsingGetRequestAddressType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBlackWhiteListsUsingGetRequestAddressType) UnmarshalJSON(b []byte) error {
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
