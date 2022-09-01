package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 用户自定义字段
type WorkitemCustomField struct {

	// 自定义字段id
	FieldId *string `json:"field_id,omitempty" xml:"field_id"`

	// 自定义字段名称
	FieldName *string `json:"field_name,omitempty" xml:"field_name"`

	// 自定义字段类型, \"Date\",\"Number\",\"DateTime\", \"MultiLineText\",\"SingleLineText\", \"Select\",  \"Checkbox\"
	FieldType *string `json:"field_type,omitempty" xml:"field_type"`

	// 自定义字段的选项源,CUSTOM，USER，DOMAIN，ITERATION，MODULE，TAG
	FieldOptionSource *string `json:"field_option_source,omitempty" xml:"field_option_source"`

	// 自定义字段值, (field_type为Date,Number,DateTime时,field_option_source为空，value值是数字的字符串)， (field_type为MultiLineText,SingleLineText时,field_option_source为空，value值是文本字符串)， (field_type为Select ,field_option_source为CUSTOM时，value值是文本字符串) (field_type为Select ,field_option_source为USER，DOMAIN，ITERATION，MODULE，TAG时，value值是Json格式{}), (field_type为Checkbox ,field_option_source为CUSTOM时，value值是字符串数组[\\\"aaa\\\"]), (field_type为\"Checkbox\" ,field_option_source为USER，DOMAIN，ITERATION，MODULE，TAG时，value值是Json的数组[{},{}])
	Value *string `json:"value,omitempty" xml:"value"`
}

func (o WorkitemCustomField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkitemCustomField struct{}"
	}

	return strings.Join([]string{"WorkitemCustomField", string(data)}, " ")
}
