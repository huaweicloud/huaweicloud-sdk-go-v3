package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"encoding/json"
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"os"
	"reflect"

	"strings"
)

type CreateNextflowJobRequestBody struct {

	// 作业的名称，取值范围：[1,63]，允许大小写字母、数字、以及特殊字符中划线(-)
	Name *def.MultiPart `json:"name"`

	// 作业的描述,取值范围：输入字符最大长度为255
	Description *def.MultiPart `json:"description,omitempty"`

	// 作业标签，取值范围[0,5]，单个标签最大长度32字符，支持中文、字母、数字、空格、下划线和中划线，且不能以空格开头或者结尾。
	Labels *def.MultiPart `json:"labels,omitempty"`

	// 作业依赖的流程id
	WorkflowId *def.MultiPart `json:"workflow_id"`

	// 流程参数列表文件，取值范围[0, 10M]
	Params *def.FilePart `json:"params,omitempty"`

	// 作业的优先级,取值范围[0,9]，0最低，默认数值0
	Priority *def.MultiPart `json:"priority,omitempty"`
}

func (o CreateNextflowJobRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNextflowJobRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNextflowJobRequestBody", string(data)}, " ")
}

func (o *CreateNextflowJobRequestBody) UnmarshalJSON(b []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	t := reflect.TypeOf(o).Elem()
	v := reflect.ValueOf(o).Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		jsonTag := t.Field(i).Tag.Get("json")
		jsonName := strings.Split(jsonTag, ",")[0]
		if m[jsonName] == nil && strings.Contains(jsonTag, "omitempty") {
			continue
		}
		field := v.FieldByName(utils.UnderscoreToCamel(jsonName))
		switch v.Field(i).Interface().(type) {
		case *def.FilePart:
			filePath := m[jsonName].(string)
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(def.NewFilePart(file)))
		case *def.MultiPart:
			field.Set(reflect.ValueOf(def.NewMultiPart(m[jsonName])))
		default:
			return errors.New(fmt.Sprintf("unmarshal %s failed", m[jsonName]))
		}
	}
	return nil
}
