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

type CreateNextflowWorkflowRequestBody struct {

	// 流程文件，文件大小[0,10M]
	WorkflowFile *def.FilePart `json:"workflow_file"`

	// 流程名称，取值范围[1,56]，允许大小写字母、数字、以及特殊字符中划线(-)和下划线(_)。更新流程时，流程名称不支持修改。
	Name *def.MultiPart `json:"name"`

	// 流程描述 取值范围[0,65535]
	Description *def.MultiPart `json:"description,omitempty"`

	// 流程标签，取值范围[0,5]，单个标签最大长度32字符，支持中文、字母、数字、空格、下划线和中划线，且不能以空格开头或者结尾。
	Labels *def.MultiPart `json:"labels,omitempty"`

	// 主文件名
	MainFile *def.MultiPart `json:"main_file"`

	// 流程参数列表文件，取值范围[0, 10M]
	Params *def.FilePart `json:"params,omitempty"`
}

func (o CreateNextflowWorkflowRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNextflowWorkflowRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNextflowWorkflowRequestBody", string(data)}, " ")
}

func (o *CreateNextflowWorkflowRequestBody) UnmarshalJSON(b []byte) error {
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
