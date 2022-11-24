package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListAddressSetListUsingGetRequest struct {

	// 防护对象id
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
