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

type CreateSecAppTaskRequestBody struct {

	// 版本：0免费版，1专业版
	Version *def.MultiPart `json:"version"`

	// 待扫描文件，后缀为.apk或.hap,专业版大小限制为2G，免费版大小限制为100M
	File *def.FilePart `json:"file"`

	// 隐私申明URL
	PrivacyStatementUrl *def.MultiPart `json:"privacy_statement_url,omitempty"`

	// 个人信息第三方共享目录URL
	PersonalInfoShareUrl *def.MultiPart `json:"personal_info_share_url,omitempty"`
}

func (o CreateSecAppTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecAppTaskRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecAppTaskRequestBody", string(data)}, " ")
}

func (o *CreateSecAppTaskRequestBody) UnmarshalJSON(b []byte) error {
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
