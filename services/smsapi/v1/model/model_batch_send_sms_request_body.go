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

type BatchSendSmsRequestBody struct {

	// 短信发送方的号码
	From *def.MultiPart `json:"from,omitempty"`

	// 短信接收方的号码
	To *def.MultiPart `json:"to,omitempty"`

	// 短信模板ID
	TemplateId *def.MultiPart `json:"templateId,omitempty"`

	// 短信模板的变量值
	TemplateParas *def.MultiPart `json:"templateParas,omitempty"`

	// SP的回调地址
	StatusCallback *def.MultiPart `json:"statusCallback,omitempty"`

	// 扩展参数，在状态报告中会原样返回。
	Extend *def.MultiPart `json:"extend,omitempty"`

	// 短信签名
	Signature *def.MultiPart `json:"signature,omitempty"`
}

func (o BatchSendSmsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSendSmsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchSendSmsRequestBody", string(data)}, " ")
}

func (o *BatchSendSmsRequestBody) UnmarshalJSON(b []byte) error {
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
