package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateBlackWhiteListDto
type UpdateBlackWhiteListDto struct {

	// 地址方向0：源地址1：目的地址
	Direction *int32 `json:"direction,omitempty"`

	// 地址类型0：ipv4,1:ipv6,2:domain
	AddressType *int32 `json:"address_type,omitempty"`

	// ip地址
	Address *string `json:"address,omitempty"`

	// 协议类型:TCP为6, UDP为17,ICMP为1,ICMPV6为58,ANY为-1
	Protocol *int32 `json:"protocol,omitempty"`

	// 端口
	Port *string `json:"port,omitempty"`

	// 黑白名单类型4：黑名单，5：白名单
	ListType *UpdateBlackWhiteListDtoListType `json:"list_type,omitempty"`

	// 防护对象id
	ObjectId *string `json:"object_id,omitempty"`
}

func (o UpdateBlackWhiteListDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBlackWhiteListDto struct{}"
	}

	return strings.Join([]string{"UpdateBlackWhiteListDto", string(data)}, " ")
}

type UpdateBlackWhiteListDtoListType struct {
	value int32
}

type UpdateBlackWhiteListDtoListTypeEnum struct {
	E_4 UpdateBlackWhiteListDtoListType
	E_5 UpdateBlackWhiteListDtoListType
}

func GetUpdateBlackWhiteListDtoListTypeEnum() UpdateBlackWhiteListDtoListTypeEnum {
	return UpdateBlackWhiteListDtoListTypeEnum{
		E_4: UpdateBlackWhiteListDtoListType{
			value: 4,
		}, E_5: UpdateBlackWhiteListDtoListType{
			value: 5,
		},
	}
}

func (c UpdateBlackWhiteListDtoListType) Value() int32 {
	return c.value
}

func (c UpdateBlackWhiteListDtoListType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateBlackWhiteListDtoListType) UnmarshalJSON(b []byte) error {
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
