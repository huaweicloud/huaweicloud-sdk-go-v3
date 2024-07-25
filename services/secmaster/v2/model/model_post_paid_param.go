package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PostPaidParam 按需订阅请求体
type PostPaidParam struct {

	// 区域ID，如cn-north-4
	RegionId string `json:"region_id"`

	// domainId
	DomainId string `json:"domain_id"`

	// 计费标签
	TagList *[]TagInfo `json:"tag_list,omitempty"`

	// 商品列表
	ProductList *[]ProductPostPaid `json:"product_list,omitempty"`

	// 操作类型：create新购, addition增加配额
	OperateType *PostPaidParamOperateType `json:"operate_type,omitempty"`
}

func (o PostPaidParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidParam struct{}"
	}

	return strings.Join([]string{"PostPaidParam", string(data)}, " ")
}

type PostPaidParamOperateType struct {
	value string
}

type PostPaidParamOperateTypeEnum struct {
	CREATE   PostPaidParamOperateType
	ADDITION PostPaidParamOperateType
}

func GetPostPaidParamOperateTypeEnum() PostPaidParamOperateTypeEnum {
	return PostPaidParamOperateTypeEnum{
		CREATE: PostPaidParamOperateType{
			value: "create",
		},
		ADDITION: PostPaidParamOperateType{
			value: "addition",
		},
	}
}

func (c PostPaidParamOperateType) Value() string {
	return c.value
}

func (c PostPaidParamOperateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPaidParamOperateType) UnmarshalJSON(b []byte) error {
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
