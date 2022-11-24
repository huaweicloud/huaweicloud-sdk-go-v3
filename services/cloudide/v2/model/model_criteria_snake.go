package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CriteriaSnake struct {

	// 过滤类型 - 1 按插件Tag过滤 - 2 按diplayName过滤 - 3 按publisherId过滤 - 4 按插件Id过滤 - 5 按插件分类过滤 - 7 按照作者名.插件名过滤 - 8 按Target（客户端）过滤 - 10 按关键字（客户端输入的）过滤 - 12 根据flags传入的值来进行过滤,eg:flags=2name就排除flags=2的插件. - 13 根据flags传入的值来进行过滤,eg:flags=2name就查询出flags=2的插件 - 18 按publisherName过滤 - 19 按publisherDisplayName过滤 - 102 按照插件状态排除插件 - 103 按照插件状态过滤出插件 - 107 supportIdeInfo - 108 根据插件ids查询
	FilterType *CriteriaSnakeFilterType `json:"filter_type,omitempty"`

	// 过滤类型对应字段名称
	Value *string `json:"value,omitempty"`
}

func (o CriteriaSnake) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CriteriaSnake struct{}"
	}

	return strings.Join([]string{"CriteriaSnake", string(data)}, " ")
}

type CriteriaSnakeFilterType struct {
	value int32
}

type CriteriaSnakeFilterTypeEnum struct {
	E_1   CriteriaSnakeFilterType
	E_2   CriteriaSnakeFilterType
	E_3   CriteriaSnakeFilterType
	E_4   CriteriaSnakeFilterType
	E_5   CriteriaSnakeFilterType
	E_7   CriteriaSnakeFilterType
	E_8   CriteriaSnakeFilterType
	E_10  CriteriaSnakeFilterType
	E_12  CriteriaSnakeFilterType
	E_13  CriteriaSnakeFilterType
	E_18  CriteriaSnakeFilterType
	E_19  CriteriaSnakeFilterType
	E_102 CriteriaSnakeFilterType
	E_103 CriteriaSnakeFilterType
	E_107 CriteriaSnakeFilterType
	E_108 CriteriaSnakeFilterType
}

func GetCriteriaSnakeFilterTypeEnum() CriteriaSnakeFilterTypeEnum {
	return CriteriaSnakeFilterTypeEnum{
		E_1: CriteriaSnakeFilterType{
			value: 1,
		}, E_2: CriteriaSnakeFilterType{
			value: 2,
		}, E_3: CriteriaSnakeFilterType{
			value: 3,
		}, E_4: CriteriaSnakeFilterType{
			value: 4,
		}, E_5: CriteriaSnakeFilterType{
			value: 5,
		}, E_7: CriteriaSnakeFilterType{
			value: 7,
		}, E_8: CriteriaSnakeFilterType{
			value: 8,
		}, E_10: CriteriaSnakeFilterType{
			value: 10,
		}, E_12: CriteriaSnakeFilterType{
			value: 12,
		}, E_13: CriteriaSnakeFilterType{
			value: 13,
		}, E_18: CriteriaSnakeFilterType{
			value: 18,
		}, E_19: CriteriaSnakeFilterType{
			value: 19,
		}, E_102: CriteriaSnakeFilterType{
			value: 102,
		}, E_103: CriteriaSnakeFilterType{
			value: 103,
		}, E_107: CriteriaSnakeFilterType{
			value: 107,
		}, E_108: CriteriaSnakeFilterType{
			value: 108,
		},
	}
}

func (c CriteriaSnakeFilterType) Value() int32 {
	return c.value
}

func (c CriteriaSnakeFilterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CriteriaSnakeFilterType) UnmarshalJSON(b []byte) error {
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
