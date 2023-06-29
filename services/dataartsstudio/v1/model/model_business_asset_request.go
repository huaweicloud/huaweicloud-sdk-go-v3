package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BusinessAssetRequest struct {

	// 关键字查询是否匹配所有属性，true为查询所有属性，false为仅查询名称描述
	SearchAllAttributes bool `json:"search_all_attributes"`

	// 标签信息 Set<String>
	Tags *interface{} `json:"tags,omitempty"`

	// 查询返回数目
	Limit int32 `json:"limit"`

	// 查询偏移量
	Offset int32 `json:"offset"`

	// 查询节点的guid
	Guid *string `json:"guid,omitempty"`

	// 查询关键字
	Query string `json:"query"`

	// 查询类型
	Type BusinessAssetRequestType `json:"type"`
}

func (o BusinessAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BusinessAssetRequest struct{}"
	}

	return strings.Join([]string{"BusinessAssetRequest", string(data)}, " ")
}

type BusinessAssetRequestType struct {
	value string
}

type BusinessAssetRequestTypeEnum struct {
	BUSINESS    BusinessAssetRequestType
	LOGICENTITY BusinessAssetRequestType
}

func GetBusinessAssetRequestTypeEnum() BusinessAssetRequestTypeEnum {
	return BusinessAssetRequestTypeEnum{
		BUSINESS: BusinessAssetRequestType{
			value: "BUSINESS",
		},
		LOGICENTITY: BusinessAssetRequestType{
			value: "LOGICENTITY",
		},
	}
}

func (c BusinessAssetRequestType) Value() string {
	return c.value
}

func (c BusinessAssetRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BusinessAssetRequestType) UnmarshalJSON(b []byte) error {
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
