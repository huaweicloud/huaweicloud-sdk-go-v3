package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DbOrTableRenameRule 库表更名规则  可以在名称添加前后缀， 填写前缀名，会给库表名称前面添加对前缀名，填写后缀名，会给库表名称后面添加对后缀名； 当前缀名和后缀名都写时，会给库表名称前后面添加对字符；
type DbOrTableRenameRule struct {

	// 前缀名称 当type为prefixAndSuffix， 填写prefix_name，库表名称仅增加前缀，若同时也填写suffix_name, 库表名称增加前后缀
	PrefixName *string `json:"prefix_name,omitempty"`

	// 后缀名称 当type为prefixAndSuffix， 填写suffix_name，库表名称仅增加后缀，若同时也填写prefix_name, 库表名称同时增加前后缀
	SuffixName *string `json:"suffix_name,omitempty"`

	// 库表映射类型 prefixAndSuffix:前缀、后缀或者前后缀
	Type *DbOrTableRenameRuleType `json:"type,omitempty"`
}

func (o DbOrTableRenameRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbOrTableRenameRule struct{}"
	}

	return strings.Join([]string{"DbOrTableRenameRule", string(data)}, " ")
}

type DbOrTableRenameRuleType struct {
	value string
}

type DbOrTableRenameRuleTypeEnum struct {
	PREFIX_AND_SUFFIX DbOrTableRenameRuleType
	MANY_TO_ONE       DbOrTableRenameRuleType
}

func GetDbOrTableRenameRuleTypeEnum() DbOrTableRenameRuleTypeEnum {
	return DbOrTableRenameRuleTypeEnum{
		PREFIX_AND_SUFFIX: DbOrTableRenameRuleType{
			value: "prefixAndSuffix",
		},
		MANY_TO_ONE: DbOrTableRenameRuleType{
			value: "manyToOne",
		},
	}
}

func (c DbOrTableRenameRuleType) Value() string {
	return c.value
}

func (c DbOrTableRenameRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DbOrTableRenameRuleType) UnmarshalJSON(b []byte) error {
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
